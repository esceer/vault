package model

import (
	"time"

	"github.com/esceer/vault/backend/internal/common"
)

type CredentialCreate struct {
	User   string
	Site   string
	Secret string
}

type CredentialResponse struct {
	ID        common.Identifier
	User      string
	Site      string
	CreatedAt time.Time
}

type SecretResponse struct {
	Secret string
}
