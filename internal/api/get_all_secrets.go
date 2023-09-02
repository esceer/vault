package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esceer/vault/internal/service"
)

type getAllSecrets struct {
	service service.VaultService
}

func NewGetAllSecretsHandler(s service.VaultService) *getAllSecrets {
	return &getAllSecrets{
		service: s,
	}
}

func (e *getAllSecrets) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostnames, err := e.service.ListKeys()
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
