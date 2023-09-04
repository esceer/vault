package setup

import (
	"github.com/esceer/vault/backend/internal/service"
	"github.com/esceer/vault/backend/internal/storage"
	"gorm.io/gorm"
)

func VaultService(db *gorm.DB) service.VaultService {
	store := storage.NewDBStore(db)
	return service.NewVaultService(store)
}
