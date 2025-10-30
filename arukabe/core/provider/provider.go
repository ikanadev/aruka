package provider

import (
	"arukabe/core/provider/handler"
	"arukabe/core/provider/repo/pg"
	"arukabe/gen/connect/provider/v1/providerv1connect"
	"arukabe/gen/sqlc"
	"context"
	"net/http"

	"github.com/anthropics/anthropic-sdk-go"
)

func RegisterProviderService(
	ctx context.Context,
	mux *http.ServeMux,
	db *sqlc.Queries,
	antClient *anthropic.Client,
) {
	repo := pg.NewPGRepository(ctx, db, antClient)
	providerHandler := handler.NewProviderHandler(repo)
	path, handler := providerv1connect.NewProviderServiceHandler(providerHandler)
	mux.Handle(path, handler)
}
