package storage

import "github.com/esceer/vault/storage/internal/inmemory"

type Store interface {
	Store(key string, secret string) error
	Retrieve(key string) (string, error)
}

func New() Store {
	return inmemory.NewStore()
}
