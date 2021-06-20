package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func get_action() (string, error) {

	log.Info("Please enter r to read, w to write, or x to exit:\n")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.Wrapf(err, "could not read line for r or w :: %s", err)
	}
	input = strings.TrimSuffix(input, "\n")

	if input == READ || input == WRITE || input == EXIT {
		return input, nil
	}
	return "", errors.Wrapf(err, "invalid action: %s, please choose r or w or x", input)
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

	var password string
	fmt.Scanf("%s", &password)

	return []byte(password), nil
}
