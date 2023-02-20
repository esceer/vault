package inmemory

type inMemoryStore struct {
	secrets map[string]string
}

func NewStore() *inMemoryStore {
	return &inMemoryStore{make(map[string]string)}
}

func (s *inMemoryStore) Store(key string, secret string) error {
	s.secrets[key] = secret
	return nil
}

func (s *inMemoryStore) Retrieve(key string) (string, error) {
	return s.secrets[key], nil
}
