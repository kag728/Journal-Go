package journal

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	clear           string = "clear"
	editorLocation  string = "entries/editor"
	textEditCommand string = "open"
)

var textEditArgs = []string{"-a", "TextEdit", editorLocation}

func GetAction() (string, error) {

	log.Info("Please enter \n\tr to read this week's entries \n\tra to read all" +
		"\n\tw to write \n\tx to exit:\n")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.Wrapf(err, "could not read line for r or w :: %s", err)
	}
	input = strings.ToLower(strings.TrimSuffix(input, "\n"))

	if input == read || input == readAll || input == write || input == exit {
		return input, nil
	}
	return "", errors.Wrapf(err, "invalid action: %s, please choose %s or %s or %s or %s", input, read, readAll, write, exit)
}

func promptForDone() error {

	log.Info("Please write entry in editor file. Press Enter when done...")

	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		return errors.Wrapf(err, "could not read new line input to signify Done")
	}
	return nil
}

func ClearScreen() {
	cmd := exec.Command(clear)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error clearing screen: %v", err)
	}
}

func openEditorInVim() error {
	cmd := exec.Command(textEditCommand, textEditArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "error opening editor in vim. Please make sure vim is installed and on the Path.")
	}
	return nil
}
