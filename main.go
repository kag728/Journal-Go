package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	action, err := get_action()
	if err != nil {
		log.Fatal(err)
	}

	run(action)
}
