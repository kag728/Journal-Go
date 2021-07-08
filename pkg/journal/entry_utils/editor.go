package entry_utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	EDITOR_FILE_NAME = "editor"
)

type Editor struct {
	editor_file_name, entry_file_name string
	decrypted_entry_contents          []byte
	current_entry, editor_file        *os.File
}

// Creates a file called "editor" in the entries/ directory.
// This is where the entry can be edited in plain text.
func CreateEditor(entry *os.File) (Editor, error) {

	editor := Editor{}
	editor.current_entry = entry
	editor.editor_file_name = path.Join(FILE_DIR, EDITOR_FILE_NAME)
	editor.entry_file_name = path.Join(editor.current_entry.Name())
	var err error

	editor.editor_file, err = os.Create(editor.editor_file_name)
	if err != nil {
		return editor, errors.Wrapf(err, "could not create editor file")
	}

	entry_contents, err := ioutil.ReadFile(editor.entry_file_name)
	if err != nil {
		return editor, errors.Wrapf(err, "error while reading contents of current entry %s", editor.entry_file_name)
	}

	decrypted_entry_contents, err := DecryptEntryContents(string(entry_contents))
	if err != nil {
		return editor, errors.Wrapf(err, "error decrypting entry contents")
	}
	editor.decrypted_entry_contents = decrypted_entry_contents

	// Only append a new line if there is already text in the file, if not we start on first line.
	var new_line string
	if string(decrypted_entry_contents) != "" {
		new_line = "\n"
	}
	editor_starting_text := fmt.Sprintf("%s%s- [%s] ", decrypted_entry_contents, new_line, time.Now().Format(time.Kitchen))
	ioutil.WriteFile(editor.editor_file_name, []byte(editor_starting_text), 7777)

	return editor, nil
}

// Encrypts the contents of the editor file and saves it to today's entry. Then deletes the editor.
func (editor *Editor) SaveEditorText() error {
	editor_contents, err := ioutil.ReadFile(editor.editor_file_name)
	if err != nil {
		return errors.Wrapf(err, "could not read editor contents")
	}
	defer editor.current_entry.Close()
	defer editor.editor_file.Close()

	editor_contents = trim_newlines(editor_contents)
	encrypted_contents, err := EncryptEditorContents(string(editor_contents))
	if err != nil {
		return errors.Wrapf(err, "error encrypting editor contents")
	}
	ioutil.WriteFile(editor.entry_file_name, encrypted_contents, 7777)

	return editor.delete()
}

func (editor *Editor) delete() error {
	err := os.Remove(editor.editor_file_name)
	if err != nil {
		return errors.Wrapf(err, "error deleting editor file")
	}
	return nil
}

func trim_newlines(editor_contents []byte) []byte {
	s_untrimmed := string(editor_contents)
	s_trimmed := strings.TrimSuffix(s_untrimmed, "\n")
	if s_trimmed != s_untrimmed {
		return trim_newlines([]byte(s_trimmed))
	}
	return []byte(s_trimmed)
}
