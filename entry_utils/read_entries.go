package entry_utils

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

func TestPassword() (bool, error) {
	entries, err := os.ReadDir(FILE_DIR)
	if err != nil {
		return false, errors.Wrapf(err, "could not open directory %s", FILE_DIR)
	}

	// If there aren't any entries, any password is good
	if len(entries) == 0 {
		return true, nil
	}

	entry := entries[0]

	entry_name := entry.Name()
	entry_contents, err := os.ReadFile(path.Join(FILE_DIR, entry_name))
	if err != nil {
		return false, errors.Wrapf(err, "error reading entry %s", entry_name)
	}

	// Check if we can decrypt with the password
	_, err = DecryptEntryContents(string(entry_contents))
	if err != nil {
		return false, nil
	}

	// Password worked
	return true, nil
}

// Print out all entries
func ReadEntries() error {

	entries, err := os.ReadDir(FILE_DIR)
	if err != nil {
		return errors.Wrapf(err, "could not open directory %s", FILE_DIR)
	}

	fmt.Printf("\n")
	for _, entry := range entries {
		entry_name := entry.Name()
		entry_name_formatted := strings.Replace(strings.Join(strings.Split(entry_name, "_")[1:], "_"), "_", " ", -1)
		entry_contents, err := os.ReadFile(path.Join(FILE_DIR, entry_name))
		if err != nil {
			return errors.Wrapf(err, "error reading entry %s", entry_name)
		}

		decrypted_entry_contents, err := DecryptEntryContents(string(entry_contents))
		if err != nil {
			return errors.Wrapf(err, "error decrypting entry contents for %s", entry_name)
		}

		fmt.Printf("%s:\n%s\n\n", entry_name_formatted, string(decrypted_entry_contents))
	}
	return nil
}
