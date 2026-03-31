package handler

import (
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	"context"

	"connectrpc.com/connect"
)

func NewChatHandler(service ChatService) *ChatHandler {
	return &ChatHandler{service}
}

type ChatHandler struct {
	service ChatService
}

// ChatMessage implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) ChatMessage(
	ctx context.Context,
	req *chatv1.ChatMessageRequest,
	stream *connect.ServerStream[chatv1.ChatMessageResponse],
) error {
	chatResult, err := c.service.ChatMessage(ctx, req)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	for {
		select {
		case result, ok := <-chatResult:
			if !ok {
				return nil
			}
			if result.Err != nil {
				return connect.NewError(connect.CodeInternal, result.Err)
			}
			resp := &chatv1.ChatMessageResponse{Delta: result.Text}
			if err := stream.Send(resp); err != nil {
				return connect.NewError(connect.CodeInternal, err)
			}
		case <-ctx.Done():
			return connect.NewError(connect.CodeCanceled, ctx.Err())
		}
	}
}

// ChatMessages implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) ChatMessages(
	ctx context.Context,
	req *chatv1.ChatMessagesRequest,
) (*chatv1.ChatMessagesResponse, error) {
	return c.service.ChatMessages(ctx, req)
}

// EditChat implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) EditChat(
	ctx context.Context,
	req *chatv1.EditChatRequest,
) (*chatv1.EditChatResponse, error) {
	return c.service.EditChat(ctx, req)
}

// ListChats implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) ListChats(
	ctx context.Context,
	req *chatv1.ListChatsRequest,
) (*chatv1.ListChatsResponse, error) {
	return c.service.ListChats(ctx, req)
}

// NewChat implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) NewChat(
	ctx context.Context,
	req *chatv1.NewChatRequest,
) (*chatv1.NewChatResponse, error) {
	return c.service.NewChat(ctx, req)
}

// AutoChatTitleUpdate implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) AutoChatTitleUpdate(
	ctx context.Context,
	req *chatv1.AutoChatTitleUpdateRequest,
) (*chatv1.AutoChatTitleUpdateResponse, error) {
	return c.service.AutoChatTitleUpdate(ctx, req)
}

// DeleteChat implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) DeleteChat(
	ctx context.Context,
	req *chatv1.DeleteChatRequest,
) (*chatv1.DeleteChatResponse, error) {
	return c.service.DeleteChat(ctx, req)
}
