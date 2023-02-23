package inmemory

type inMemoryStore struct {
	secrets map[string][]byte
}

func NewStore() *inMemoryStore {
	return &inMemoryStore{make(map[string][]byte)}
}

func (s *inMemoryStore) Store(key string, secret []byte) error {
	s.secrets[key] = secret
	return nil
}

func (s *inMemoryStore) Retrieve(key string) ([]byte, error) {
	return s.secrets[key], nil
}

func (s *inMemoryStore) Delete(key string) error {
	delete(s.secrets, key)
	return nil
}
