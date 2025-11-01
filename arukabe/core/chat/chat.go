package chat

import (
	"arukabe/core/chat/handler"
	"arukabe/core/chat/repo/pg"
	"arukabe/gen/connect/chat/v1/chatv1connect"
	"arukabe/gen/sqlc"
	"context"
	"net/http"

	"github.com/anthropics/anthropic-sdk-go"
)

func RegisterChatService(
	ctx context.Context,
	mux *http.ServeMux,
	db *sqlc.Queries,
	antClient *anthropic.Client,
) {
	repo := pg.NewPGRepository(ctx, db, antClient)
	chatHandler := handler.NewChatHandler(repo)
	path, handler := chatv1connect.NewChatServiceHandler(chatHandler)
	mux.Handle(path, handler)
}
