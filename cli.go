package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func get_action() (string, error) {

	log.Info("Please enter r or w to read or write:\n")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read line for r or w :: %s", err)
	}
	input = strings.TrimSuffix(input, "\n")

	if input == READ || input == WRITE {
		return input, nil
	}
	return "", fmt.Errorf("invalid action: %s, please choose r or w", input)
}

func prompt_for_done() error {

	log.Info("Please write entry in editor file. Press Enter when done...")

	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("could not read new line input to signify Done :: %s", err)
	}
	return nil
}
