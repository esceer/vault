package setup

import (
	"github.com/esceer/vault/internal/api"
	"github.com/esceer/vault/internal/middleware"
	"github.com/esceer/vault/internal/service"
	"github.com/gorilla/mux"
)

func WebRouting(vaultService service.VaultService) *mux.Router {
	getAllSecretsHandler := api.NewGetAllSecretsHandler(vaultService)

	middleware := middleware.NewLoggingMiddleware()

	router := mux.NewRouter()
	router.Handle("/password", middleware.WithHTTPLogging(getAllSecretsHandler))
	return router
}
