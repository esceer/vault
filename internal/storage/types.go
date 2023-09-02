package storage

import "time"

type Credential struct {
	ID        int
	User      string
	Key       string
	Secret    []byte
	CreatedAt time.Time
}
