package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Printf("Please enter r or w to read or write:\n")
	action, err := get_action()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(action)
	_, err = get_current_entry()
	if err != nil {
		log.Fatalf("Could not create or open today's entry :: %s", err)
	}
}

func get_action() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Could not read line for r or w.")
	}
	input = strings.TrimSuffix(input, "\n")

	if input == "r" || input == "w" {
		return input, nil
	}
	return "", fmt.Errorf("invalid action :: %s, please choose r or w", input)
}
