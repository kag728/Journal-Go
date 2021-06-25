package entry_utils

import (
	"io/fs"
	"strconv"
	"strings"
)

// Filter list of entries so each on begins with a number and ends with a number.
// Definitely not a great implementation but we'll see how it does
func filter_entries(entries []fs.DirEntry) []fs.DirEntry {

	filtered_entries := []fs.DirEntry{}
	for _, entry := range entries {

		entry_name_split := strings.Split(entry.Name(), "_")

		add_to_list := true
		_, err := strconv.Atoi(entry_name_split[0])
		if err != nil {
			add_to_list = false
		}

		_, err = strconv.Atoi(entry_name_split[len(entry_name_split)-1])
		if err != nil {
			add_to_list = false
		}

		if add_to_list {
			filtered_entries = append(filtered_entries, entry)
		}
	}

	return filtered_entries
}
