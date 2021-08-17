package authentication

import (
	"journal/pkg/journal/entry_utils"
	"os"
	"path"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/term"
)

// File containing your passphrase
const (
	passphrase_file_name = ".internal/.passphrase"
	passphrase_length    = 32
)

func Authenticate(encryptor *Encryptor) {

	password, err := os.ReadFile(passphrase_file_name)
	if err != nil {
		log.Warnf("Could not find passphrase file at %s", passphrase_file_name)
		var prompt_err error
		password, prompt_err = prompt_for_password()
		if prompt_err != nil {
			log.Fatal(errors.Wrapf(err, "error getting password"))
		}
	} else {

		// If the full password is present
		if len(password) == passphrase_length {
			log.Info("Obtained password from file.")

			// If not, user needs to give the remaining bytes
		} else {
			remaining_bytes := passphrase_length - len(password)
			log.Infof("Please enter %d remaining bytes:", remaining_bytes)

			pin, err := read_password()
			if err != nil {
				log.Fatal(errors.Wrapf(err, "error reading pin"))
			}

			pin_bytes := []byte(pin)
			old_password := password
			password = append(old_password, pin_bytes...)
		}
	}

	encryptor.SetPassword(password)

	authenticated, err := TestPassword(encryptor)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "error authenticating"))
	}

	if !authenticated {
		log.Fatal("Password incorrect.")
	}
}

func prompt_for_password() ([]byte, error) {

	log.Info("Please enter password:")
	return read_password()
}

func read_password() ([]byte, error) {
	byte_password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error reading password from input")
	}

	return byte_password, nil
}

func TestPassword(encryptor *Encryptor) (bool, error) {

	// If the entries folder doesn't exist, user can go ahead and create one
	entries, err := os.ReadDir(entry_utils.FILE_DIR)
	if err != nil {
		log.Warn("Entries folder does not exist, create a new entry to set password.")
		return true, nil
	}
	entries = entry_utils.Filter_entries(entries)

	// If there aren't any entries, any password is good
	if len(entries) == 0 {
		return true, nil
	}

	entry := entries[0]
	entry_name := entry.Name()
	entry_contents, err := os.ReadFile(path.Join(entry_utils.FILE_DIR, entry_name))
	if err != nil {
		return false, errors.Wrapf(err, "error reading entry %s", entry_name)
	}

	// Check if we can decrypt with the password
	_, err = encryptor.DecryptEntryContents(string(entry_contents))
	if err != nil {
		return false, errors.Wrapf(err, "error decrypting with the password")
	}

	// Password worked
	return true, nil
}
