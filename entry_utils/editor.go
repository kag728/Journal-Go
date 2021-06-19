package entry_utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

const (
	EDITOR_FILE_NAME = "editor"
)

func Create_editor(current_entry *os.File) (*os.File, error) {

	file, err := os.Create(path.Join(FILE_DIR, EDITOR_FILE_NAME))
	if err != nil {
		return file, fmt.Errorf("could not create editor file :: %s", err)
	}
	entry_contents, err := ioutil.ReadFile(path.Join(FILE_DIR, current_entry.Name()))
	if err != nil {
		return file, fmt.Errorf("error while reading contents of current entry %s :: %s", current_entry.Name(), err)
	}
	log.Infof("Contents of editor: %s", string(entry_contents))
	ioutil.WriteFile(path.Join(FILE_DIR, EDITOR_FILE_NAME), entry_contents, 7777)
	return file, nil
}
