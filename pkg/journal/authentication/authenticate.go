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
	passphraseFileName = ".internal/.passphrase"
	passphraseLength   = 32
)

func Authenticate(encryptor *Encryptor) {

	password, err := os.ReadFile(passphraseFileName)
	if err != nil {
		log.Warnf("Could not find passphrase file at %s", passphraseFileName)
		var promptErr error
		password, promptErr = promptForPassword()
		if promptErr != nil {
			log.Fatal(errors.Wrapf(err, "error getting password"))
		}
	} else {

		// If the full password is present
		if len(password) == passphraseLength {
			log.Info("Obtained password from file.")

			// If not, user needs to give the remaining bytes
		} else {
			remainingBytes := passphraseLength - len(password)
			log.Infof("Please enter %d remaining bytes:", remainingBytes)

			pin, err := readPassword()
			if err != nil {
				log.Fatal(errors.Wrapf(err, "error reading pin"))
			}

			pinBytes := []byte(pin)
			oldPassword := password
			password = append(oldPassword, pinBytes...)
		}
	}

	encryptor.SetPassword(password)

	authenticated, err := testPassword(encryptor)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "error authenticating"))
	}

	if !authenticated {
		log.Fatal("Password incorrect.")
	}
}

func promptForPassword() ([]byte, error) {

	log.Info("Please enter password:")
	return readPassword()
}

func readPassword() ([]byte, error) {
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error reading password from input")
	}

	return bytePassword, nil
}

func testPassword(encryptor *Encryptor) (bool, error) {

	// If the entries folder doesn't exist, user can go ahead and create one
	entries, err := os.ReadDir(entry_utils.FileDir)
	if err != nil {
		log.Warn("Entries folder does not exist, create a new entry to set password.")
		return true, nil
	}
	entries = entry_utils.FilterEntries(entries)

	// If there aren't any entries, any password is good
	if len(entries) == 0 {
		return true, nil
	}

	entry := entries[0]
	entryName := entry.Name()
	entryContents, err := os.ReadFile(path.Join(entry_utils.FileDir, entryName))
	if err != nil {
		return false, errors.Wrapf(err, "error reading entry %s", entryName)
	}

	// Check if we can decrypt with the password
	_, err = encryptor.DecryptEntryContents(string(entryContents))
	if err != nil {
		return false, errors.Wrapf(err, "error decrypting with the password")
	}

	// Password worked
	return true, nil
}
