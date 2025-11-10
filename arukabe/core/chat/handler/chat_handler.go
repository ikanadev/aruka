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
func (c *ChatHandler) ChatMessage(context.Context, *connect.Request[chatv1.ChatMessageRequest], *connect.ServerStream[chatv1.ChatMessageResponse]) error {
	panic("unimplemented")
}

// ChatMessages implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) ChatMessages(context.Context, *connect.Request[chatv1.ChatMessagesRequest]) (*connect.Response[chatv1.ChatMessagesResponse], error) {
	panic("unimplemented")
}

// EditChat implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) EditChat(context.Context, *connect.Request[chatv1.EditChatRequest]) (*connect.Response[chatv1.EditChatResponse], error) {
	panic("unimplemented")
}

// ListChats implements chatv1connect.ChatServiceHandler.
func (c *ChatHandler) ListChats(context.Context, *connect.Request[chatv1.ListChatsRequest]) (*connect.Response[chatv1.ListChatsResponse], error) {
	panic("unimplemented")
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
