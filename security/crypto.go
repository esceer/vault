package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func Encode(key []byte, text []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Galois/Counter Mode, is a mode of operation for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	// New byte array with the size of the nonce which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// Populates nonce with a cryptographically secure random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Seal encrypts and authenticates text, authenticates the additional data
	// and appends the encodedText to dst (first parameter), returning the updated slice.
	// The nonce must be NonceSize() bytes long and unique for all time, for a given key.
	return gcm.Seal(nonce, nonce, text, nil), nil
}

func Decode(key []byte, cipherText []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("Ciphertext is too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}
