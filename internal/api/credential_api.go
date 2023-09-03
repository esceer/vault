package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esceer/vault/internal/api/model"
	"github.com/esceer/vault/internal/common"
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

func (a *credentialApi) GetAll(w http.ResponseWriter, r *http.Request) {
	creds, err := a.service.GetAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(creds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *credentialApi) GetSecret(w http.ResponseWriter, r *http.Request) {
	id, err := getIntPathParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	masterkey, err := getQuerySecretParam(r, "masterkey")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	secret, err := a.service.GetSecret(common.Identifier(id), common.Secret(masterkey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = json.NewEncoder(w).Encode(secret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *credentialApi) Save(w http.ResponseWriter, r *http.Request) {
	masterkey, err := getQuerySecretParam(r, "masterkey")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cred, err := getBody[model.CredentialCreate](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = a.service.Save(&cred, common.Secret(masterkey)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (a *credentialApi) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getIntPathParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = a.service.Delete(common.Identifier(id)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
