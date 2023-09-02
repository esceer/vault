package setup

import (
	"github.com/esceer/vault/internal/service"
	"github.com/esceer/vault/internal/storage"
	"gorm.io/gorm"
)

func VaultService(db *gorm.DB) service.VaultService {
	store := storage.NewDBStore(db)
	return service.NewVaultService(store)
}
