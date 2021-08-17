package entries_io

import (
	"fmt"
	"journal/pkg/journal/authentication"
	"journal/pkg/journal/entry_utils"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Print out all entries
func ReadEntries(encryptor *authentication.Encryptor, one_week bool) error {

	entries, err := os.ReadDir(entry_utils.FILE_DIR)
	if err != nil {
		log.Warn("The entries directory does not exist, please create an entry first.")
		return nil
	}

	if one_week {
		entries, err = entry_utils.Filter_entries_for_week(entries)
		if err != nil {
			return errors.Wrapf(err, "error filtering entries for this week")
		}
	} else {
		entries = entry_utils.Filter_entries(entries)
		if err != nil {
			return errors.Wrapf(err, "error filtering entries for this week")
		}
	}

	fmt.Printf("\n")
	for _, entry := range entries {
		entry_name := entry.Name()

		// Get every part of file name aside from prefix
		entry_sections := strings.Split(entry_name, "_")[1:]

		// If it starts with a weekday, add a comma
		if len(entry_sections) == entry_utils.Max_entry_name_sections-1 {
			entry_sections[0] = fmt.Sprintf("%s,", entry_sections[0])
		}

		entry_name_formatted := strings.Join(entry_sections, " ")
		entry_contents, err := os.ReadFile(path.Join(entry_utils.FILE_DIR, entry_name))
		if err != nil {
			return errors.Wrapf(err, "error reading entry %s", entry_name)
		}

		decrypted_entry_contents, err := encryptor.DecryptEntryContents(string(entry_contents))
		if err != nil {
			return errors.Wrapf(err, "error decrypting entry contents for %s", entry_name)
		}

		fmt.Printf("%s:\n%s\n\n", entry_name_formatted, string(decrypted_entry_contents))
	}
	return nil
}
