package anthropicmsg

import (
	"arukabe/core/common/mappers"
	"arukabe/core/common/types"
	"arukabe/gen/sqlc"
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/packages/param"
	"github.com/google/uuid"
)

func HandleAnthropicChatMessage(
	ctx context.Context,
	chat sqlc.Chat,
	model sqlc.Model,
	messages []types.Message, // Last message is the user message
	antClient *anthropic.Client,
	db *sqlc.Queries,
) (<-chan types.ChatStreamResult, error) {
	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages provided")
	}
	anthropicMessages, err := messagesToAnthropicMessages(messages)
	if err != nil {
		return nil, err
	}
	messageParams := anthropic.MessageNewParams{
		MaxTokens:   8192,
		Messages:    anthropicMessages,
		Model:       anthropic.Model(model.ModelIdentifier),
		Temperature: param.Opt[float64]{Value: 0.3},
	}
	if len(chat.Prompt) > 0 {
		messageParams.System = []anthropic.TextBlockParam{{Text: chat.Prompt}}
	}
	stream := antClient.Messages.NewStreaming(ctx, messageParams)

	resultChan := make(chan types.ChatStreamResult)

	anthropicResponse := anthropic.Message{}
	go func() {
		defer close(resultChan)

		for stream.Next() {
			event := stream.Current()
			err := anthropicResponse.Accumulate(event)
			if err != nil {
				select {
				case resultChan <- types.ChatStreamResult{Err: err}:
				case <-ctx.Done():
				}
				return
			}
			switch eventVariant := event.AsAny().(type) {
			case anthropic.ContentBlockDeltaEvent:
				switch deltaVariant := eventVariant.Delta.AsAny().(type) {
				case anthropic.TextDelta:
					select {
					case resultChan <- types.ChatStreamResult{Text: deltaVariant.Text}:
					case <-ctx.Done():
						return
					}
					// TODO: handle other delta types
				}
			default:
				continue
			}
		}
		if err := stream.Err(); err != nil {
			select {
			case resultChan <- types.ChatStreamResult{Err: err}:
			case <-ctx.Done():
			}
			return
		}
		err = saveUserMessageAndAnthropicResponse(ctx, chat.ID, messages[len(messages)-1], anthropicResponse, db)
		if err != nil {
			select {
			case resultChan <- types.ChatStreamResult{Err: err}:
			case <-ctx.Done():
			}
			return
		}
	}()

	return resultChan, nil
}

func saveUserMessageAndAnthropicResponse(
	ctx context.Context,
	chatID uuid.UUID,
	userMsg types.Message,
	response anthropic.Message,
	db *sqlc.Queries,
) error {
	var responseSections types.MessageSections
	err := responseSections.FromAnthropicMessageSections(response.Content)
	if err != nil {
		return err
	}

	responseMessage := types.Message{
		Content: responseSections,
	}

	userMsgID, err := uuid.NewV7()
	if err != nil {
		return err
	}
	responseMsgID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	userMsgJSON, err := userMsg.Content.MarshalJSON()
	if err != nil {
		return err
	}
	responseMsgJSON, err := responseMessage.Content.MarshalJSON()
	if err != nil {
		return err
	}

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

	_, err = db.SaveMessages(ctx, toSave)
	return err
}

func messagesToAnthropicMessages(messages []types.Message) ([]anthropic.MessageParam, error) {
	anthropicMessages := make([]anthropic.MessageParam, len(messages))
	for i, msg := range messages {
		role, err := mappers.RoleToAnthropicRoleParam(msg.Role)
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
