package provider

import (
	"arukabe/core/provider/handler"
	"arukabe/core/provider/repository"
	"arukabe/core/provider/service"
	"arukabe/gen/connect/provider/v1/providerv1connect"
	"arukabe/gen/sqlc"
	"net/http"

	"github.com/anthropics/anthropic-sdk-go"
)

func RegisterProviderService(
	mux *http.ServeMux,
	db *sqlc.Queries,
	antClient *anthropic.Client,
) {
	repo := repository.NewProviderRepository(db, antClient)
	service := service.NewProviderService(repo)
	providerHandler := handler.NewProviderHandler(service)
	path, handler := providerv1connect.NewProviderServiceHandler(providerHandler)
	mux.Handle(path, handler)
}
