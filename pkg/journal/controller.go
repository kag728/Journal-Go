package journal

import (
	"journal/pkg/journal/authentication"
	"journal/pkg/journal/entries_io"
	"journal/pkg/uploader"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	write   = "w"
	read    = "r"
	readAll = "ra"
	exit    = "x"
)

func Run(action string, encryptor *authentication.Encryptor) {

	switch action {
	case write:
		handleWrite(encryptor)
	case read:
		handleRead(encryptor)
	case readAll:
		handleReadAll(encryptor)
	case exit:
		handleExit()
	default:
		log.Warn("Could not interpret input")
	}
}

func handleWrite(encryptor *authentication.Encryptor) {
	entry, err := entries_io.GetCurrentEntry()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error getting current entry"))
	}

	editor, err := authentication.CreateEditor(entry, encryptor)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error creating editor"))
	}

	err = openEditorInVim()
	if err != nil {
		log.Warnf("Could not open vim. Please open %s with another text editor.", editorLocation)
	}

	err = promptForDone()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error prompting for done"))
	}

	err = editor.SaveEditorText(encryptor)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error saving editor text"))
	}

	log.Info("Encrypted contents of editor and saved them to entry.")

	log.Info("Uploading changes to cloud directory...")

	uploadName, err := uploader.Upload()
	if err != nil {

		if _, ok := err.(*uploader.CloudConfigNotFound); ok {
			log.Warn("Skipping cloud upload, could not find cloudconfig file.")
		} else {
			log.Error(errors.Wrapf(err, "Entry was not uploaded to cloud folder. Please make sure folder name is correct"))
		}

	} else {
		log.Infof("Successfully uploaded %s", uploadName)
	}
}

func handleRead(encryptor *authentication.Encryptor) {
	err := entries_io.ReadEntries(encryptor, true)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "error reading entries"))
	}
}

func handleReadAll(encryptor *authentication.Encryptor) {
	err := entries_io.ReadEntries(encryptor, false)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "error reading entries"))
	}
}

func handleExit() {
	ClearScreen()
	os.Exit(0)
}
