package main

import (
	"journal/entry_utils"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// File containing your passphrase
const PASSPHRASE_FILE = ".internal/.passphrase"

func authenticate() {

	var password []byte
	password, err := os.ReadFile(PASSPHRASE_FILE)
	if err != nil {
		log.Warn("Could not find passphrase file at %s", PASSPHRASE_FILE)
		var prompt_err error
		password, prompt_err = prompt_for_password()
		if prompt_err != nil {
			log.Fatal(errors.Wrapf(err, "error getting password"))
		}
	} else {

		// If the full password is present
		if len(password) == 32 {
			log.Info("obtained password from file.")

			// If not, user needs to give the remaining bytes
		} else {
			remaining_bytes := 32 - len(password)
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

	entry_utils.SetPassword(password)

	authenticated, err := entry_utils.TestPassword()
	if err != nil {
		log.Fatal(errors.Wrapf(err, "error authenticating"))
	}

	if !authenticated {
		log.Fatal("Password incorrect.")
	}
}
