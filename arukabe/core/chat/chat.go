package chat

import (
	"arukabe/core/chat/handler"
	"arukabe/core/chat/repository"
	"arukabe/core/chat/service"
	"arukabe/gen/connect/chat/v1/chatv1connect"
	"arukabe/gen/sqlc"
	"net/http"

	"github.com/anthropics/anthropic-sdk-go"
)

func RegisterChatService(
	mux *http.ServeMux,
	db *sqlc.Queries,
	antClient *anthropic.Client,
) {
	repo := repository.NewChatRepository(db, antClient)
	service := service.NewChatService(repo)
	chatHandler := handler.NewChatHandler(*service)
	path, handler := chatv1connect.NewChatServiceHandler(chatHandler)
	mux.Handle(path, handler)
}
