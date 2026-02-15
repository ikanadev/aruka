package service

import (
	"arukabe/core/common/constants"
	"arukabe/core/common/mappers"
	"arukabe/core/common/types"
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	"arukabe/gen/sqlc"
	"context"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/packages/param"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (cs *ChatService) AutoChatTitleUpdate(
	ctx context.Context,
	req *chatv1.AutoChatTitleUpdateRequest,
) (*chatv1.AutoChatTitleUpdateResponse, error) {
	chatUUID, err := uuid.Parse(req.ChatId)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	dbMessages, err := cs.db.GetChatMessages(ctx, chatUUID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if len(dbMessages) == 0 {
		return &chatv1.AutoChatTitleUpdateResponse{Title: ""}, nil
	}

	rawMessages := make([]sqlc.Message, len(dbMessages))
	for i, row := range dbMessages {
		rawMessages[i] = row.Message
	}

	messages, err := mappers.FromDBMessages(rawMessages)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	conversationText := buildConversationText(messages)

	title, err := generateTitle(ctx, cs.antClient, conversationText)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal,
			fmt.Errorf("failed to generate title: %w", err))
	}

	err = cs.db.UpdateChat(ctx, sqlc.UpdateChatParams{
		ChatID: chatUUID,
		Title:  pgtype.Text{String: title, Valid: true},
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return &chatv1.AutoChatTitleUpdateResponse{Title: title}, nil
}

func buildConversationText(messages []types.Message) string {
	maxMessagesLength := 1500
	var sb strings.Builder
	for _, msg := range messages {
		role := "User"
		if msg.Role == sqlc.MessageRoleASSISTANT {
			role = "Assistant"
		}
		if sb.Len() > maxMessagesLength {
			break
		}
		for _, section := range msg.Content {
			if section.Type == types.MessageSectionTypeText {
				textContent := section.Payload.(types.MessageTextSection)
				sb.WriteString(role)
				sb.WriteString(": ")
				sb.WriteString(textContent.Text)
				sb.WriteString("\n")
			}
		}
	}
	return sb.String()
}

func generateTitle(ctx context.Context, client *anthropic.Client, conversationText string) (string, error) {
	resp, err := client.Messages.New(ctx, anthropic.MessageNewParams{
		Model:       constants.TitleGenerationModel,
		MaxTokens:   50,
		Temperature: param.Opt[float64]{Value: 0.2},
		System: []anthropic.TextBlockParam{
			{Text: "Generate a short, descriptive title (max 6 words) for the following conversation. Reply with only the title, nothing else."},
		},
		Messages: []anthropic.MessageParam{
			{
				Role: anthropic.MessageParamRoleUser,
				Content: []anthropic.ContentBlockParamUnion{
					{OfText: &anthropic.TextBlockParam{Text: conversationText}},
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	for _, block := range resp.Content {
		switch v := block.AsAny().(type) {
		case anthropic.TextBlock:
			return strings.TrimSpace(v.Text), nil
		}
	}

	return "", fmt.Errorf("no text content in title generation response")
}
