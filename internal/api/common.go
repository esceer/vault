package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/esceer/vault/internal/common"
	"github.com/gorilla/mux"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func getQueryParam(r *http.Request, param string) (string, error) {
	qparams := r.URL.Query()
	p, ok := qparams[param]
	if !ok {
		return "", fmt.Errorf("missing parameter: %v", param)
	}
	return p[0], nil
}

func getQuerySecretParam(r *http.Request, param string) (common.Secret, error) {
	p, err := getQueryParam(r, param)
	if err != nil {
		return nil, err
	}
	return base64.URLEncoding.DecodeString(p)
}

func getPathParam(r *http.Request, param string) (string, error) {
	vars := mux.Vars(r)
	p, ok := vars[param]
	if !ok {
		return "", fmt.Errorf("missing parameter: %v", param)
	}
	return p, nil
}

func getIntPathParam(r *http.Request, param string) (int, error) {
	p, err := getPathParam(r, param)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(p)
}

func getBody[T any](r *http.Request) (T, error) {
	var t T
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	return t, d.Decode(&t)
}
