package repository

import (
	"arukabe/core/common/mappers"
	"arukabe/core/common/types"
	"arukabe/gen/sqlc"
	"context"
	"errors"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/packages/param"
	"github.com/google/uuid"
)

func (cr *ChatRepository) HandleAnthropicChatMessage(
	ctx context.Context,
	chat sqlc.Chat,
	model sqlc.Model,
	// Last message is the user message
	messages []types.Message,
) (<-chan string, <-chan error, error) {
	if len(messages) == 0 {
		return nil, nil, errors.New("no messages provided")
	}
	anthropicMessages, err := messagesToAnthropicMessages(messages)
	if err != nil {
		return nil, nil, err
	}
	stream := cr.antClient.Messages.NewStreaming(ctx, anthropic.MessageNewParams{
		// TODO: make this configurable
		MaxTokens:   8192,
		Messages:    anthropicMessages,
		Model:       anthropic.Model(model.ModelIdentifier),
		Temperature: param.Opt[float64]{Value: 0.3},
		System:      []anthropic.TextBlockParam{{Text: chat.Prompt}},
	})

	textChan := make(chan string)
	errChan := make(chan error)

	anthropicResponse := anthropic.Message{}
	go func() {
		defer close(textChan)
		defer close(errChan)

		for stream.Next() {
			event := stream.Current()
			err := anthropicResponse.Accumulate(event)
			if err != nil {
				errChan <- err
				break
			}
			switch eventVariant := event.AsAny().(type) {
			case anthropic.ContentBlockDeltaEvent:
				switch deltaVariant := eventVariant.Delta.AsAny().(type) {
				case anthropic.TextDelta:
					select {
					case textChan <- deltaVariant.Text:
					case <-ctx.Done():
						errChan <- ctx.Err()
						return
					}
					// TODO: handle other delta types
				}
			default:
				continue
			}
		}
		if err := stream.Err(); err != nil {
			errChan <- err
		}
		go saveUserMessageAndAnthropicResponse(cr.db, chat.ID, messages[len(messages)-1], anthropicResponse)
	}()

	return textChan, errChan, nil
}

func saveUserMessageAndAnthropicResponse(db *sqlc.Queries, chatID uuid.UUID, userMsg types.Message, response anthropic.Message) {
	ctx := context.Background()
	var responseSections types.MessageSections
	_ = responseSections.FromAnthropicMessageSections(response.Content)

	responseMessage := types.Message{
		Content: responseSections,
	}

	userMsgID, _ := uuid.NewV7()
	responseMsgID, _ := uuid.NewV7()

	userMsgJSON, _ := userMsg.Content.MarshalJSON()
	responseMsgJSON, _ := responseMessage.Content.MarshalJSON()

	toSave := []sqlc.SaveMessagesParams{
		{
			ID:      userMsgID,
			ChatID:  chatID,
			Role:    sqlc.MessageRoleUSER,
			Content: userMsgJSON,
		},
		{
			ID:      responseMsgID,
			ChatID:  chatID,
			Role:    sqlc.MessageRoleASSISTANT,
			Content: responseMsgJSON,
		},
	}

	db.SaveMessages(ctx, toSave)
}

func messagesToAnthropicMessages(messages []types.Message) ([]anthropic.MessageParam, error) {
	anthropicMessages := make([]anthropic.MessageParam, len(messages))
	for i, msg := range messages {
		role, err := mappers.RoleToAnthropic(msg.Role)
		if err != nil {
			return nil, err
		}
		content, err := msg.Content.ToAnthropicMessageSections()
		if err != nil {
			return nil, err
		}
		anthropicMessages[i] = anthropic.MessageParam{
			Content: content,
			Role:    role,
		}
	}
	return anthropicMessages, nil
}
