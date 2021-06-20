package main

import (
	"journal/entry_utils"

	log "github.com/sirupsen/logrus"
)

const (
	WRITE = "w"
	READ  = "r"
)

func run(action string) {

	if action == WRITE {
		entry, err := entry_utils.GetCurrentEntry()
		if err != nil {
			log.Fatal(err)
		}

		_, err = entry_utils.CreateEditor(entry)
		if err != nil {
			log.Fatal(err)
		}

		err = prompt_for_done()
		if err != nil {
			log.Fatal(err)
		}

		err = entry_utils.SaveEditorText()
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Saved contents of editor to entry.")

	} else {
		log.Fatal("Unsupported action.")
	}
}
