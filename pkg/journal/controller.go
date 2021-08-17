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
	write    = "w"
	read     = "r"
	read_all = "ra"
	exit     = "x"
)

func Run(action string, encryptor *authentication.Encryptor) {

	if action == write {
		entry, err := entries_io.GetCurrentEntry()
		if err != nil {
			log.Fatal(errors.Wrap(err, "error getting current entry"))
		}

		editor, err := authentication.CreateEditor(entry, encryptor)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error creating editor"))
		}

		err = open_editor_in_vim()
		if err != nil {
			log.Warnf("Could not open vim. Please open %s with another text editor.", vim_location)
		}

		err = prompt_for_done()
		if err != nil {
			log.Fatal(errors.Wrap(err, "error prompting for done"))
		}

		err = editor.SaveEditorText(encryptor)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error saving editor text"))
		}

		log.Info("Encrypted contents of editor and saved them to entry.")

		log.Info("Uploading changes to cloud directory...")

		upload_name, err := uploader.Upload()
		if err != nil {

			if _, ok := err.(*uploader.CloudConfigNotFound); ok {
				log.Warn("Skipping cloud upload, could not find cloudconfig file.")
			} else {
				log.Error(errors.Wrapf(err, "Entry was not uploaded to cloud folder. Please make sure folder name is correct"))
			}

		} else {
			log.Infof("Successfully uploaded %s", upload_name)
		}

	} else if action == read {

		err := entries_io.ReadEntries(encryptor, true)
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error reading entries"))
		}

	} else if action == read_all {

		err := entries_io.ReadEntries(encryptor, false)
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error reading entries"))
		}

	} else if action == exit {
		ClearScreen()
		os.Exit(0)
	} else {
		log.Warn("Could not interpret input")
	}
}
