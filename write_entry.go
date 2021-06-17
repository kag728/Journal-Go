package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const FILE_DIR string = "./entries"

func get_current_entry() (*os.File, error) {
	_, dir_err := os.Stat(FILE_DIR)
	if dir_err != nil {
		mkdir_err := os.Mkdir(FILE_DIR, 0777)
		if mkdir_err != nil {
			log.Errorf("could not create dir %s :: %s", FILE_DIR, mkdir_err)
		}
	}

	entry_name, err := get_entry_name()
	if err != nil {
		log.Debug("Found file for today, stopping search.")
	}

	file_name := path.Join(FILE_DIR, entry_name)
	file, open_err := os.Open(file_name)
	if open_err != nil {
		file, create_err := os.Create(file_name)
		if create_err != nil {
			log.Errorf("Could not open or create %s :: %s", file_name, create_err)
			return file, create_err
		}
		log.Infof("Created file %s", file_name)
		return file, nil
	}
	log.Infof("today's entry already exists, opening")
	return file, nil
}

func get_entry_name() (string, error) {

	entries, err := os.ReadDir(FILE_DIR)
	if err != nil {
		log.Errorf("Could not open directory")
		return "", nil
	}

	time := time.Now()
	entry_date := fmt.Sprintf("%s_%d_%d", time.Month(), time.Day(), time.Year())

	prefix := 0
	for _, entry := range entries {
		entry_name := entry.Name()
		entry_name = strings.Join(strings.Split(entry_name, "_")[1:], "_")

		if entry_name == entry_date {
			return "", fmt.Errorf("%s already exists", entry_date)
		}
		prefix++
	}

	entry_name := fmt.Sprintf("%d_%s", prefix, entry_date)
	return entry_name, nil
}
