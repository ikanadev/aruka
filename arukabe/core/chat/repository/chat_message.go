package repository

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/packages/param"
)

func (cr *ChatRepository) ChatMessage(ctx context.Context) (<-chan string, <-chan error) {
	stream := cr.antClient.Messages.NewStreaming(ctx, anthropic.MessageNewParams{
		MaxTokens: 8192,
		Messages: []anthropic.MessageParam{
			{
				Content: []anthropic.ContentBlockParamUnion{
					{
						OfText: &anthropic.TextBlockParam{
							Text: "Hello I'm Dave, greet me!",
						},
					},
				},
				Role: anthropic.MessageParamRoleUser,
			},
		},
		Model:       anthropic.ModelClaudeHaiku4_5,
		Temperature: param.Opt[float64]{Value: 0.3},
		System: []anthropic.TextBlockParam{
			{
				Text: "You're a futuristic human shape robot assistant. You're too realistic that you have to somehow clarify you're a robot at the start of any conversation but not in boring or obvious ways, always a creative way.",
			},
		},
	})
	textChan := make(chan string)
	errChan := make(chan error)

	// message := anthropic.Message{}
	go func() {
		defer close(textChan)
		defer close(errChan)

		for stream.Next() {
			event := stream.Current()
			// err := message.Accumulate(event)

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
				}
			default:
				continue
			}
		}
		if err := stream.Err(); err != nil {
			errChan <- err
		}
	}()

	return textChan, errChan
}
