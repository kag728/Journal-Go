package entry_utils

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/errors"
)

const (
	EDITOR_FILE_NAME = "editor"
)

var (
	editor_file_name, entry_file_name string
	current_entry                     *os.File
	editor                            *os.File
)

func CreateEditor(entry *os.File) (*os.File, error) {

	current_entry = entry
	editor_file_name = path.Join(FILE_DIR, EDITOR_FILE_NAME)
	entry_file_name = path.Join(current_entry.Name())

	var err error
	editor, err = os.Create(editor_file_name)
	if err != nil {
		return editor, errors.Wrapf(err, "could not create editor file")
	}

	entry_contents, err := ioutil.ReadFile(entry_file_name)
	if err != nil {
		return editor, errors.Wrapf(err, "error while reading contents of current entry %s", entry_file_name)
	}

	ioutil.WriteFile(editor_file_name, entry_contents, 7777)
	return editor, nil
}

func SaveEditorText() error {
	editor_contents, err := ioutil.ReadFile(editor_file_name)
	if err != nil {
		return errors.Wrapf(err, "could not read editor contents")
	}
	defer current_entry.Close()
	defer editor.Close()

	ioutil.WriteFile(entry_file_name, editor_contents, 7777)
	return nil
}
