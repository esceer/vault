package storage

type identifier int

type Store interface {
	GetKeysByUser(user string) ([]string, error)
	GetSecret(id identifier) ([]byte, error)
	Store(user, key string, secret []byte) error
	Delete(id identifier) error
}
