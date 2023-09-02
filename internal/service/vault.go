package service

import "github.com/esceer/vault/internal/storage"

type VaultService interface {
	ListKeys() ([]string, error)
}

type vaultService struct {
	store storage.Store
}

func NewVaultService(s storage.Store) VaultService {
	return &vaultService{s}
}

func (s vaultService) ListKeys() ([]string, error) {
	return s.store.GetKeysByUser(CurrentUser())
}
