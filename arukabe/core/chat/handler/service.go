package handler

import (
	"arukabe/core/common/types"
	chatv1 "arukabe/gen/connect/aruka/chat/v1"
	"context"
)

type ChatService interface {
	NewChat(ctx context.Context, req *chatv1.NewChatRequest) (*chatv1.NewChatResponse, error)
	ChatMessage(ctx context.Context, req *chatv1.ChatMessageRequest) (<-chan types.ChatStreamResult, error)
	ChatMessages(ctx context.Context, req *chatv1.ChatMessagesRequest) (*chatv1.ChatMessagesResponse, error)
	EditChat(ctx context.Context, req *chatv1.EditChatRequest) (*chatv1.EditChatResponse, error)
	ListChats(ctx context.Context, req *chatv1.ListChatsRequest) (*chatv1.ListChatsResponse, error)
	AutoChatTitleUpdate(ctx context.Context, req *chatv1.AutoChatTitleUpdateRequest) (*chatv1.AutoChatTitleUpdateResponse, error)
	DeleteChat(context.Context, *chatv1.DeleteChatRequest) (*chatv1.DeleteChatResponse, error)
}
