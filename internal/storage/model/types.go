package model

import (
	"time"

	"github.com/esceer/vault/internal/common"
)

type Credential struct {
	ID        common.Identifier
	User      string
	Site      string
	Secret    common.Secret
	CreatedAt time.Time
}
