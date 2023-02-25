package storage

import "github.com/esceer/vault/storage/internal/db"

type IStore interface {
	ListKeys() ([]string, error)
	Retrieve(key string) ([]byte, error)
	Store(key string, secret []byte) error
	Delete(key string) error
}

func New() IStore {
	return db.NewStore()
}
