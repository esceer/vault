package main

import (
	"net/http"

	"github.com/esceer/vault/internal/storage"
	"github.com/esceer/vault/internal/web"
	"github.com/gorilla/mux"
)

func main() {
	store := storage.New()
	router := mux.NewRouter()
	web.SetupRouting(router, store)
	http.ListenAndServe(":8080", router)
}
