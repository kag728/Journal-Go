package journal

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/term"
)

func GetAction() (string, error) {

	log.Info("Please enter r to read, w to write, or x to exit:\n")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.Wrapf(err, "could not read line for r or w :: %s", err)
	}
	input = strings.ToLower(strings.TrimSuffix(input, "\n"))

	if input == READ || input == WRITE || input == EXIT {
		return input, nil
	}
	return "", errors.Wrapf(err, "invalid action: %s, please choose %s or %s or %s", input, READ, WRITE, EXIT)
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

func prompt_for_password() ([]byte, error) {

	log.Info("Please enter password:")
	return read_password()
}

func read_password() ([]byte, error) {
	byte_password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return []byte{}, errors.Wrapf(err, "error reading password from input")
	}

	return byte_password, nil
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
