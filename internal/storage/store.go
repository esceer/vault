package storage

type identifier int

type Store interface {
	GetKeysByUser(string) ([]string, error)
	GetById(identifier) (*Credential, error)
	Save(*Credential) error
	Delete(identifier) error
}
