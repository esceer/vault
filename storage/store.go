package storage

import "github.com/esceer/vault/storage/internal/inmemory"

type IStore interface {
	Retrieve(key string) ([]byte, error)
	Store(key string, secret []byte) error
	Delete(key string) error
}

func New() IStore {
	return inmemory.NewStore()
}
