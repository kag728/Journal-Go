package main

import (
	"fmt"
	"journal/entry_utils"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	WRITE = "w"
	READ  = "r"
	EXIT  = "x"
)

const PASSPHRASE_FILE = ".internal/.passphrase"

func authenticate() {

	var password []byte
	password, err := os.ReadFile(PASSPHRASE_FILE)
	if err != nil {
		log.Warn("Could not find passphrase file at .internal/.passphrase")
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

			var pin string
			fmt.Scanf("%s", &pin)

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

func run(action string) {

	if action == WRITE {
		entry, err := entry_utils.GetCurrentEntry()
		if err != nil {
			log.Fatal(errors.Wrap(err, "error getting current entry"))
		}

		_, err = entry_utils.CreateEditor(entry)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error creating editor"))
		}

		err = prompt_for_done()
		if err != nil {
			log.Fatal(errors.Wrap(err, "error prompting for done"))
		}

		err = entry_utils.SaveEditorText()
		if err != nil {
			log.Fatal(errors.Wrap(err, "error saving editor text"))
		}

		log.Infof("Saved contents of editor to entry.")

	} else if action == READ {

		err := entry_utils.ReadEntries()
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error reading entries"))
		}

	} else if action == EXIT {
		os.Exit(0)
	} else {
		log.Warn("Could not interpret input")
	}
}
