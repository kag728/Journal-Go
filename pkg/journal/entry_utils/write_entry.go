package entry_utils

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// The directory containing the entries
const FILE_DIR string = "entries"

// Either opens today's entry if it's already been made, or creates a new one and returns that.
// The file is named based on the current date
func GetCurrentEntry() (*os.File, error) {
	_, dir_err := os.Stat(FILE_DIR)
	if dir_err != nil {
		mkdir_err := os.Mkdir(FILE_DIR, 0777)
		if mkdir_err != nil {
			return &os.File{}, errors.Wrapf(mkdir_err, "could not create dir %s", FILE_DIR)
		}
	}

	entry_name, err := get_entry_name()
	if err != nil {
		return &os.File{}, errors.Wrapf(err, "error getting file name")
	}

	file_name := path.Join(FILE_DIR, entry_name)
	file, open_err := os.Open(file_name)
	if open_err != nil {
		file, create_err := os.Create(file_name)
		if create_err != nil {
			return file, errors.Wrapf(create_err, "could not open or create %s", file_name)
		}

		log.Infof("Created file %s", file_name)
		return file, nil
	}

	return file, nil
}

func get_entry_name() (string, error) {

	entries, err := os.ReadDir(FILE_DIR)
	if err != nil {
		return "", errors.Wrapf(err, "could not open directory %s", FILE_DIR)
	}
	entries = filter_entries(entries)

	time := time.Now()
	entry_date := fmt.Sprintf("%s_%s_%d_%d", time.Weekday(), time.Month(), time.Day(), time.Year())

	prefix := 0
	found_today_entry := false
	for _, entry := range entries {

		entry_name := entry.Name()
		if entry_name != "editor" {

			entry_prefix, err := strconv.Atoi(strings.Split(entry_name, "_")[0])
			if err != nil {
				return "", errors.Wrapf(err, "error converting prefix %d to an integer", entry_prefix)
			}
			prefix = entry_prefix
			entry_name = strings.Join(strings.Split(entry_name, "_")[1:], "_")

			if entry_name == entry_date {
				found_today_entry = true
				break
			}
		}
	}

	if !found_today_entry {
		prefix = len(entries)
	}

	formatted_prefix, err := fill_prefix(prefix)
	if err != nil {
		return "", errors.Wrapf(err, "error formatting prefix")
	}
	entry_name := fmt.Sprintf("%s_%s", formatted_prefix, entry_date)

	return entry_name, nil
}
