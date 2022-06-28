package authentication

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"journal/pkg/journal/entry_utils"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	EditorFileName = "editor"
)

type Editor struct {
	editorFileName, entryFileName string
	decryptedEntryContents        []byte
	currentEntry, editorFile      *os.File
}

// CreateEditor creates a file called "editor" in the entries/ directory.
// This is where the entry can be edited in plain text.
func CreateEditor(entry *os.File, encryptor *Encryptor) (Editor, error) {

	editor := Editor{}
	editor.currentEntry = entry
	editor.editorFileName = path.Join(entry_utils.FileDir, EditorFileName)
	editor.entryFileName = path.Join(editor.currentEntry.Name())
	var err error

	editor.editorFile, err = os.Create(editor.editorFileName)
	if err != nil {
		return editor, errors.Wrapf(err, "could not create editor file")
	}

	entryContents, err := ioutil.ReadFile(editor.entryFileName)
	if err != nil {
		return editor, errors.Wrapf(err, "error while reading contents of current entry %s", editor.entryFileName)
	}

	decryptedEntryContents, err := encryptor.DecryptEntryContents(string(entryContents))
	if err != nil {
		return editor, errors.Wrapf(err, "error decrypting entry contents")
	}
	editor.decryptedEntryContents = decryptedEntryContents

	// Only append a new line if there is already text in the file, if not we start on first line.
	var newLine string
	if string(decryptedEntryContents) != "" {
		newLine = "\n"
	}
	editorStartingText := fmt.Sprintf("%s%s- [%s] ", decryptedEntryContents, newLine, time.Now().Format(time.Kitchen))
	err = ioutil.WriteFile(editor.editorFileName, []byte(editorStartingText), 7777)
	if err != nil {
		return editor, errors.Wrapf(err, "could not write starting text to editor file")
	}

	return editor, nil
}

// SaveEditorText encrypts the contents of the editor file and saves it to today's entry. Then deletes the editor.
func (editor *Editor) SaveEditorText(encryptor *Encryptor) error {
	editorContents, err := ioutil.ReadFile(editor.editorFileName)
	if err != nil {
		return errors.Wrapf(err, "could not read editor contents")
	}
	defer func(currentEntry *os.File) {
		err := currentEntry.Close()
		if err != nil {
			log.Errorf("could not close entry %s", currentEntry.Name())
		}
	}(editor.currentEntry)
	defer func(editorFile *os.File) {
		err := editorFile.Close()
		if err != nil {
			log.Errorf("could not close editor file")
		}
	}(editor.editorFile)

	editorContents = trimNewlines(editorContents)
	encryptedContents, err := encryptor.EncryptEditorContents(string(editorContents))
	if err != nil {
		return errors.Wrapf(err, "error encrypting editor contents")
	}
	err = ioutil.WriteFile(editor.entryFileName, encryptedContents, 7777)
	if err != nil {
		return errors.Wrapf(err, "could not write encrypted file to %s", editor.entryFileName)
	}

	return editor.delete()
}

func (editor *Editor) delete() error {
	err := os.Remove(editor.editorFileName)
	if err != nil {
		return errors.Wrapf(err, "error deleting editor file")
	}
	return nil
}

func trimNewlines(editorContents []byte) []byte {
	sUntrimmed := string(editorContents)
	sTrimmed := strings.TrimSuffix(sUntrimmed, "\n")
	if sTrimmed != sUntrimmed {
		return trimNewlines([]byte(sTrimmed))
	}
	return []byte(sTrimmed)
}
