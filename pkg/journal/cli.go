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
	clear             string = "clear"
	vim               string = "vim"
	editor_location   string = "entries/editor"
	text_edit_command string = "open"
)

var text_edit_args []string = []string{"-a", "TextEdit", editor_location}

func GetAction() (string, error) {

	log.Info("Please enter r to read this week's, ra to read all, w to write, or x to exit:\n")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.Wrapf(err, "could not read line for r or w :: %s", err)
	}
	input = strings.ToLower(strings.TrimSuffix(input, "\n"))

	if input == read || input == read_all || input == write || input == exit {
		return input, nil
	}
	return "", errors.Wrapf(err, "invalid action: %s, please choose %s or %s or %s or %s", input, read, read_all, write, exit)
}

func prompt_for_done() error {

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
	cmd.Run()
}

func open_editor_in_vim() error {
	cmd := exec.Command(text_edit_command, text_edit_args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "error opening editor in vim. Please make sure vim is installed and on the Path.")
	}
	return nil
}
