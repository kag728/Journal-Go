package main

import (
	"journal/pkg/journal"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {

	journal.Authenticate()

	for {
		action, err := journal.GetAction()
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error getting action"))
		}

		journal.Run(action)
	}
}
