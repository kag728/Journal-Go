package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func get_action() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read line for r or w.")
	}
	input = strings.TrimSuffix(input, "\n")

	if input == READ || input == WRITE {
		return input, nil
	}
	return "", fmt.Errorf("invalid action :: %s, please choose r or w", input)
}
