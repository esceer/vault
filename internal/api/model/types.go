package model

import (
	"time"

	"github.com/esceer/vault/internal/common"
)

type CredentialCreate struct {
	User   string
	Site   string
	Secret common.Secret
}

type CredentialResponse struct {
	ID        common.Identifier
	User      string
	Site      string
	CreatedAt time.Time
}

type SecretResponse struct {
	Secret common.Secret
}
