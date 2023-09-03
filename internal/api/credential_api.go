package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esceer/vault/internal/service"
)

type credentialApi struct {
	service service.VaultService
}

func NewCredentialApiHandler(s service.VaultService) *credentialApi {
	return &credentialApi{
		service: s,
	}
}

func (e *credentialApi) GetAll(w http.ResponseWriter, r *http.Request) {
	creds, err := e.service.GetAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(creds)
}

func (e *credentialApi) GetSecret(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (e *credentialApi) Save(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (e *credentialApi) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO
}
