package main

import (
	"journal/pkg"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {

	pkg.Authenticate()

	for {
		action, err := pkg.GetAction()
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error getting action"))
		}

		pkg.Run(action)
	}
}
