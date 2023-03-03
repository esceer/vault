package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esceer/vault/internal/storage"
)

type GetAllSecrets struct {
	store storage.IStore
}

func NewEngine(store storage.IStore) *GetAllSecrets {
	return &GetAllSecrets{
		store: store,
	}
}

func (e *GetAllSecrets) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostnames, err := e.store.ListKeys()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	response, err := json.Marshal(&HostnamesResponse{hostnames})
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Write(response)
}
