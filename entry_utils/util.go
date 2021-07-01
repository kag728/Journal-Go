package entry_utils

import (
	"fmt"
	"io/fs"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const max_prefix_length = 4

// Filter list of entries so each on begins with a number and ends with a number.
// Definitely not a great implementation but we'll see how it does
func filter_entries(entries []fs.DirEntry) []fs.DirEntry {

	filtered_entries := []fs.DirEntry{}
	for _, entry := range entries {

		entry_name_split := strings.Split(entry.Name(), "_")

		if (len(entry_name_split)) != 4 {
			continue
		}

		_, err := strconv.Atoi(entry_name_split[0])
		if err != nil {
			continue
		}

		_, err = strconv.Atoi(entry_name_split[len(entry_name_split)-1])
		if err != nil {
			continue
		}

		filtered_entries = append(filtered_entries, entry)
	}

	return filtered_entries
}

func get_prefix_digits(n int) int {
	digits := 0
	for i := n; i >= 1; i = i / 10 {
		digits++
	}
	return digits
}

func fill_prefix(n int) (string, error) {
	digits := get_prefix_digits(n)
	if digits > max_prefix_length {
		return "", errors.Wrapf(errors.New("Prefix error"), "Entry prefix too high. Max possible prefix: %d", max_prefix_length)
	}

	output := ""
	for i := 0; i < max_prefix_length-digits; i++ {
		output += "0"
	}

	return fmt.Sprintf("%s%d", output, n), nil
}
