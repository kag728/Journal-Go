package uploader

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"journal/pkg/journal/entries_io"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

const cloudConfig = ".internal/cloudconfig"

type CloudConfigNotFound struct {
	filename, err string
}

func (f *CloudConfigNotFound) Error() string {
	return fmt.Sprintf("Could not open file: %s due to error: %s", f.filename, f.err)
}

// Upload uploads today's entry to the folder specified in the cloud_config file.
func Upload() (string, error) {

	cloudDir, err := getCloudDir()
	if err != nil {
		return "", &CloudConfigNotFound{filename: cloudConfig}
	}

	entry, err := entries_io.GetCurrentEntry()
	if err != nil {
		return "", errors.Wrapf(err, "error getting current entry")
	}

	entryContents, err := ioutil.ReadFile(entry.Name())
	if err != nil {
		return "", errors.Wrapf(err, "could not read editor contents")
	}
	defer func(entry *os.File) {
		err := entry.Close()
		if err != nil {
			log.Errorf("could not close entry %s", entry.Name())
		}
	}(entry)

	entryName := path.Join(
		cloudDir,
		strings.Split(entry.Name(), "/")[1],
	)
	_, err = os.Create(entryName)
	if err != nil {
		return "", errors.Wrapf(err, "error creating entry in cloud folder")
	}

	err = ioutil.WriteFile(entryName, entryContents, 7777)
	if err != nil {
		return "", errors.Wrapf(err, "error writing contents to cloud entry")
	}

	return entryName, nil
}

func getCloudDir() (string, error) {
	cloudDir, err := os.ReadFile(cloudConfig)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(cloudDir), "\n"), nil
}
