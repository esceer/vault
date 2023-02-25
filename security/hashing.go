package security

import (
	"golang.org/x/crypto/argon2"
)

var (
	hashingParams *argon2Parameters
)

type argon2Parameters struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func init() {
	hashingParams = &argon2Parameters{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
}

func Hash32(secret, salt []byte) []byte {
	return argon2.IDKey(secret, salt, hashingParams.iterations, hashingParams.memory, hashingParams.parallelism, hashingParams.keyLength)
}
