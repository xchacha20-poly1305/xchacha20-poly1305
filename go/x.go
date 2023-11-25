package xchacha20poly1305

import (
	"crypto/cipher"

	"golang.org/x/crypto/chacha20poly1305"
)

func New(key []byte) (cipher.AEAD, error) {
	return chacha20poly1305.NewX(key)
}
