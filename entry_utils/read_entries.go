package entry_utils

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Test the supplied password on the first file in the entries folder
func TestPassword() (bool, error) {

	// If the entries folder doesn't exist, user can go ahead and create one
	entries, err := os.ReadDir(FILE_DIR)
	if err != nil {
		log.Warn("Entries folder does not exist, create a new entry to set password.")
		return true, nil
	}
	entries = filter_entries(entries)

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
		return false, errors.Wrapf(err, "error decrypting with the password")
	}

	// Password worked
	return true, nil
}

// Print out all entries
func ReadEntries() error {

	entries, err := os.ReadDir(FILE_DIR)
	if err != nil {
		log.Warn("The entries directory does not exist, please create an entry first.")
		return nil
	}
	entries = filter_entries(entries)

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
