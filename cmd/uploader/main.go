package main

import (
	"journal/pkg/uploader"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Uploading entries to cloud.")

	entry_name, err := uploader.Upload()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("Successfully uploaded %s", entry_name)
}
