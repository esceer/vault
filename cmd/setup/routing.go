package setup

import (
	"github.com/esceer/vault/internal/api"
	"github.com/esceer/vault/internal/middleware"
	"github.com/esceer/vault/internal/service"
	"github.com/gorilla/mux"
)

func WebRouting(vaultService service.VaultService) *mux.Router {
	middleware := middleware.NewLoggingMiddleware()
	credentialApi := api.NewCredentialApiHandler(vaultService)

	router := mux.NewRouter()
	vr := router.PathPrefix("/vault").Subrouter()

	cr := vr.PathPrefix("/credentials").Subrouter()
	cr.HandleFunc("", middleware.WithHTTPLogging(credentialApi.GetAll)).Methods("GET")
	cr.HandleFunc("", middleware.WithHTTPLogging(credentialApi.Save)).Methods("POST")
	cr.HandleFunc("/{id}", middleware.WithHTTPLogging(credentialApi.Delete)).Methods("DELETE")
	cr.HandleFunc("/{id}/secret", middleware.WithHTTPLogging(credentialApi.GetSecret)).Methods("GET")
	return router
}
