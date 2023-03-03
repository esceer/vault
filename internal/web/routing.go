package web

import (
	"fmt"

	"github.com/esceer/vault/internal/storage"
	"github.com/gorilla/mux"
)

func SetupRouting(router *mux.Router, store storage.IStore) {
	engine := &GetAllSecrets{store}
	fmt.Println("Listening...")
	router.Handle("/password", engine)
}
