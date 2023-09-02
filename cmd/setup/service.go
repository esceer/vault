package setup

import (
	"database/sql"

	"github.com/esceer/vault/internal/service"
	"github.com/esceer/vault/internal/storage"
)

func VaultService(db *sql.DB) service.VaultService {
	store := storage.NewDBStorage(db)
	return service.NewVaultService(store)
}
