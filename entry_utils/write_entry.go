package entry_utils

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const FILE_DIR string = "entries1"

func GetCurrentEntry() (*os.File, error) {
	_, dir_err := os.Stat(FILE_DIR)
	if dir_err != nil {
		mkdir_err := os.Mkdir(FILE_DIR, 7777)
		if mkdir_err != nil {
			log.Errorf("could not create dir %s :: %s", FILE_DIR, mkdir_err)
		}
	}

	entry_name, err := get_entry_name()
	if err != nil {
		log.Debug("Found file for today, stopping search.")
	}

	file_name := path.Join(entry_name)
	file, open_err := os.Open(file_name)
	if open_err != nil {
		file, create_err := os.Create(file_name)
		if create_err != nil {
			log.Errorf("Could not open or create %s :: %s", file_name, create_err)
			return file, create_err
		}
		log.Debugf("Created file %s", file_name)
		return file, nil
	}
	log.Debugf("Today's entry already exists, opening %s", path.Join(FILE_DIR, file.Name()))
	return file, nil
}

func get_entry_name() (string, error) {

	entries, err := os.ReadDir(FILE_DIR)
	if err != nil {
		log.Errorf("Could not open directory")
		return "", fmt.Errorf("could not open directory %s :: %s", FILE_DIR, err)
	}

	time := time.Now()
	entry_date := fmt.Sprintf("%s_%d_%d", time.Month(), time.Day(), time.Year())

	prefix := 0
	for _, entry := range entries {
		entry_name := entry.Name()
		entry_name = strings.Join(strings.Split(entry_name, "_")[1:], "_")

		if entry_name == entry_date {
			break
		}
		prefix++
	}

	entry_name := fmt.Sprintf("%d_%s", prefix, entry_date)
	return entry_name, nil
}
