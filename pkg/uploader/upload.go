package uploader

import (
	"fmt"
	"io/ioutil"
	"journal/pkg/journal/entry_utils"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

const cloud_config = ".internal/cloudconfig"

type CloudConfigNotFound struct {
	filename, err string
}

func (f *CloudConfigNotFound) Error() string {
	return fmt.Sprintf("Could not open file: %s due to error: %s", f.filename, f.err)
}

// Uploads today's entry to the folder specified in the cloud_config file.
func Upload() (string, error) {

	cloud_dir, err := get_cloud_dir()
	if err != nil {
		return "", &CloudConfigNotFound{filename: cloud_config}
	}

	entry, err := entry_utils.GetCurrentEntry()
	if err != nil {
		return "", errors.Wrapf(err, "error getting current entry")
	}

	entry_contents, err := ioutil.ReadFile(entry.Name())
	if err != nil {
		return "", errors.Wrapf(err, "could not read editor contents")
	}
	defer entry.Close()

	entry_name := path.Join(
		cloud_dir,
		strings.Split(entry.Name(), "/")[1],
	)
	_, err = os.Create(entry_name)
	if err != nil {
		return "", errors.Wrapf(err, "error creating entry in cloud folder")
	}

	err = ioutil.WriteFile(entry_name, entry_contents, 7777)
	if err != nil {
		return "", errors.Wrapf(err, "error writing contents to cloud entry")
	}

	return entry_name, nil
}

func get_cloud_dir() (string, error) {
	cloud_dir, err := os.ReadFile(cloud_config)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(cloud_dir), "\n"), nil
}
