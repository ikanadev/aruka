package chat

import (
	"arukabe/core/chat/handler"
	"arukabe/core/chat/service"
	"arukabe/gen/connect/aruka/chat/v1/chatv1connect"
	"arukabe/gen/sqlc"
	"net/http"

	"github.com/anthropics/anthropic-sdk-go"
)

func RegisterChatService(
	mux *http.ServeMux,
	db *sqlc.Queries,
	antClient *anthropic.Client,
) {
	service := service.NewChatService(db, antClient)
	chatHandler := handler.NewChatHandler(service)
	path, handler := chatv1connect.NewChatServiceHandler(chatHandler)
	mux.Handle(path, handler)
}
