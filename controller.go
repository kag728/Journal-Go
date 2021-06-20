package main

import (
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

func authenticate() {
	password, err := prompt_for_password()
	if err != nil {
		log.Fatal(errors.Wrapf(err, "error getting password"))
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