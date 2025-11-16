package grpc

import (
	"arukabe/core/chat"
	"arukabe/core/provider"
	"arukabe/db"
	"arukabe/gen/sqlc"
	"arukabe/utils"
	"context"
	"net/http"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupGRPCServer() {
	config := utils.GetConfig()
	ctx := context.Background()

	// Database connection
	pool, err := pgxpool.New(ctx, config.DBConn)
	utils.PanicIfErr(err)
	defer pool.Close()

	// sqlc
	queries := sqlc.New(pool)

	// Database setup
	utils.PanicIfErr(db.MigrateDB(config))
	utils.PanicIfErr(db.SeedProviders(ctx, queries))

	// Anthropic
	antClient := anthropic.NewClient(option.WithAPIKey(config.AnthropicAPIKey))

	// Handler
	mux := http.NewServeMux()
	handler := chain(mux, corsMiddleware)
	provider.RegisterProviderService(mux, queries, &antClient)
  chat.RegisterChatService(mux, queries, &antClient)


	p := new(http.Protocols)
	p.SetHTTP1(true)
	p.SetUnencryptedHTTP2(true)
	httpServer := http.Server{
		Addr:     "0.0.0.0:" + config.GrpcPort,
		Handler:  handler,
		Protocols: p,
	}
	utils.PanicIfErr(httpServer.ListenAndServe())
}
