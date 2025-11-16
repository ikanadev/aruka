package handler

import (
	"arukabe/core/chat/service"
	chatv1 "arukabe/gen/connect/chat/v1"
	"context"

	"connectrpc.com/connect"
)

func NewChatHandler(service service.ChatService) *ChatHandler {
	return &ChatHandler{service}
}

type ChatHandler struct {
	service service.ChatService
}

// ChatMessage implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) ChatMessage(
	ctx context.Context,
	req *connect.Request[chatv1.ChatMessageRequest],
	stream *connect.ServerStream[chatv1.ChatMessageResponse],
) error {
	chatResult, err := c.service.ChatMessage(ctx, req.Msg)
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
	req *connect.Request[chatv1.ChatMessagesRequest],
) (*connect.Response[chatv1.ChatMessagesResponse], error) {
	resp, err := c.service.ChatMessages(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// EditChat implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) EditChat(
	ctx context.Context,
	req *connect.Request[chatv1.EditChatRequest],
) (*connect.Response[chatv1.EditChatResponse], error) {
	resp, err := c.service.EditChat(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// ListChats implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) ListChats(
	ctx context.Context,
	req *connect.Request[chatv1.ListChatsRequest],
) (*connect.Response[chatv1.ListChatsResponse], error) {
	resp, err := c.service.ListChats(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

// NewChat implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) NewChat(
	ctx context.Context,
	req *connect.Request[chatv1.NewChatRequest],
) (*connect.Response[chatv1.NewChatResponse], error) {
	resp, err := c.service.NewChat(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
