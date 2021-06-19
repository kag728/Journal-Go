package main

import (
	"fmt"
	"journal/entry_utils"

	log "github.com/sirupsen/logrus"
)

const (
	WRITE = "w"
	READ  = "r"
)

func main() {
	fmt.Printf("Please enter r or w to read or write:\n")
	action, err := get_action()
	if err != nil {
		log.Fatal(err)
	}

	if action == WRITE {
		entry, err := entry_utils.Get_current_entry()
		if err != nil {
			log.Fatalf("Could not create or open today's entry :: %s", err)
		}

		log.Infof("Creating editor for entry %s...", entry.Name())
		_, err = entry_utils.Create_editor(entry)
		if err != nil {
			log.Fatalf("Error creating editor :: %s", err)
		}

	} else {
		log.Fatal("Unsupported action.")
	}

}
