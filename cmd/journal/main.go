package main

import (
	"journal/pkg/journal"
	"journal/pkg/journal/authentication"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {

	encryptor := &authentication.Encryptor{}

	authentication.Authenticate(encryptor)

	for {
		action, err := journal.GetAction()
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error getting action"))
		}

		journal.Run(action, encryptor)
	}
}
