package pkg

import (
	"journal/pkg/journal/entry_utils"
	"journal/pkg/uploader"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	WRITE = "w"
	READ  = "r"
	EXIT  = "x"
)

func Run(action string) {

	if action == WRITE {
		entry, err := entry_utils.GetCurrentEntry()
		if err != nil {
			log.Fatal(errors.Wrap(err, "error getting current entry"))
		}

		editor, err := entry_utils.CreateEditor(entry)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error creating editor"))
		}

		err = prompt_for_done()
		if err != nil {
			log.Fatal(errors.Wrap(err, "error prompting for done"))
		}

		err = editor.SaveEditorText()
		if err != nil {
			log.Fatal(errors.Wrap(err, "error saving editor text"))
		}

		log.Infof("Saved contents of editor to entry.")

		log.Info("Uploading changes to cloud directory...")
		upload_name, err := uploader.Upload()
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error uploading changes"))
		}
		log.Info("Successfully uploaded %s", upload_name)

	} else if action == READ {

		err := entry_utils.ReadEntries()
		if err != nil {
			log.Fatal(errors.Wrapf(err, "error reading entries"))
		}

	} else if action == EXIT {
		ClearScreen()
		os.Exit(0)
	} else {
		log.Warn("Could not interpret input")
	}
}
