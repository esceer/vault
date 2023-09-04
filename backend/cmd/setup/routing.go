package setup

import (
	"net/http"

	apispec "github.com/esceer/vault/backend/api"
	"github.com/esceer/vault/backend/internal/api"
	"github.com/esceer/vault/backend/internal/middleware"
	"github.com/esceer/vault/backend/internal/service"
	swaggerui "github.com/esceer/vault/swagger-ui"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func WebRouting(vaultService service.VaultService) {
	router := mux.NewRouter()
	vr := router.PathPrefix("/vault").Subrouter()

	cr := vr.PathPrefix("/credentials").Subrouter()
	credentialApi := api.NewCredentialApiHandler(vaultService)
	cr.HandleFunc("", credentialApi.GetAll).Methods("GET")
	cr.HandleFunc("", credentialApi.Save).Methods("POST")
	cr.HandleFunc("/{id}", credentialApi.Delete).Methods("DELETE")
	cr.HandleFunc("/{id}/secret", credentialApi.GetSecret).Methods("GET")

	http.Handle("/", withCorsSetup(withLogging(router)))
	http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", swaggerui.Handler(apispec.ApiSpec)))
}

func withLogging(h http.Handler) http.Handler {
	middleware := middleware.NewLoggingMiddleware()
	return middleware.WithHTTPLogging(h)
}

func withCorsSetup(h http.Handler) http.Handler {
	return cors.AllowAll().Handler(h)
}
