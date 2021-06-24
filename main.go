package main

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {

	authenticate()

	for {
		action, err := get_action()
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error getting action"))
		}

		run(action)
	}
}
