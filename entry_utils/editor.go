package entry_utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

const (
	EDITOR_FILE_NAME = "editor"
)

var (
	editor_file_name, entry_file_name string
	current_entry                     *os.File
	editor                            *os.File
)

func CreateEditor(entry *os.File) (*os.File, error) {

	current_entry = entry
	editor_file_name = path.Join(FILE_DIR, EDITOR_FILE_NAME)
	entry_file_name = path.Join(current_entry.Name())

	var err error
	editor, err = os.Create(editor_file_name)
	if err != nil {
		return editor, errors.Wrapf(err, "could not create editor file")
	}

	entry_contents, err := ioutil.ReadFile(entry_file_name)
	if err != nil {
		return editor, errors.Wrapf(err, "error while reading contents of current entry %s", entry_file_name)
	}

	decrypted_entry_contents, err := decrypt_entry_content(string(entry_contents))
	if err != nil {
		return editor, errors.Wrapf(err, "error decrypting entry contents")
	}

	ioutil.WriteFile(editor_file_name, decrypted_entry_contents, 7777)
	return editor, nil
}

func SaveEditorText() error {
	editor_contents, err := ioutil.ReadFile(editor_file_name)
	if err != nil {
		return errors.Wrapf(err, "could not read editor contents")
	}
	defer current_entry.Close()
	defer editor.Close()

	encrypted_contents, err := encrypt_editor_contents(string(editor_contents))
	if err != nil {
		return errors.Wrapf(err, "error encrypting editor contents")
	}
	ioutil.WriteFile(entry_file_name, encrypted_contents, 7777)

	return delete_editor()
}

func delete_editor() error {
	err := os.Remove(editor_file_name)
	if err != nil {
		return errors.Wrapf(err, "error deleting editor file")
	}
	return nil
}

func encrypt_editor_contents(contents string) ([]byte, error) {
	text := []byte(contents)
	key := []byte("passphrasewhichneedstobe32bytes!")

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	// if there are any errors, handle them
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

func decrypt_entry_content(contents string) ([]byte, error) {

	contents_bytes := []byte(contents)
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error creating new cipher with key")
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error creating decryption gcm")
	}

	nonceSize := gcm.NonceSize()
	if len(contents_bytes) < nonceSize {
		return []byte{}, errors.Wrapf(err, "the length of the entry contents is less than nonce size")
	}

	nonce, ciphertext := contents_bytes[:nonceSize], contents_bytes[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error getting plain text")
	}

	return plaintext, nil
}
