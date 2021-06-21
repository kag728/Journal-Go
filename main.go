package main

import (
	"log"

	"github.com/pkg/errors"
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
