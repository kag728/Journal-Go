package authentication

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
)

// Encryptor struct for encrypting and decrypting strings of text using a provided key
type Encryptor struct {
	key []byte
}

// EncryptEditorContents encrypts the contents currently saved in the editor file and returns them
func (e *Encryptor) EncryptEditorContents(contents string) ([]byte, error) {
	text := []byte(contents)

	c, err := aes.NewCipher(e.key)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error while creating cipher")
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error while creating gcm")
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, errors.Wrapf(err, "error getting random numbers into nonce")
	}

	ret := gcm.Seal(nonce, nonce, text, nil)
	return ret, nil
}

// DecryptEntryContents decrypts and returns the contents of today's entry
func (e *Encryptor) DecryptEntryContents(contents string) ([]byte, error) {
	contentsBytes := []byte(contents)

	c, err := aes.NewCipher(e.key)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error creating new cipher with key")
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error creating decryption gcm")
	}

	nonceSize := gcm.NonceSize()
	if len(contentsBytes) < nonceSize {
		return []byte{}, errors.Wrapf(err, "the length of the entry contents is less than nonce size")
	}

	nonce, ciphertext := contentsBytes[:nonceSize], contentsBytes[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error getting plain text")
	}

	return plaintext, nil
}

func (e *Encryptor) SetPassword(p []byte) {
	e.key = p
}
