package entry_utils

import (
	"fmt"
	"io/fs"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const max_prefix_length = 4

const max_entry_name_sections = 5
const min_entry_name_sections = 4

// Filter list of entries so each on begins with a number and ends with a number.
// Definitely not a great implementation but we'll see how it does
func filter_entries(entries []fs.DirEntry) []fs.DirEntry {

	filtered_entries := []fs.DirEntry{}
	for _, entry := range entries {

		entry_name_split := strings.Split(entry.Name(), "_")

		if (len(entry_name_split)) > max_entry_name_sections || (len(entry_name_split)) < min_entry_name_sections {
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

func filter_entries_for_week(entries []fs.DirEntry) ([]fs.DirEntry, error) {

	entries = filter_entries(entries)

	now := time.Now()
	today := now.Weekday()

	days_back := get_days_back(fmt.Sprint(today))
	hours_back := time.Duration(days_back * 24)

	start_date := now.Add(-hours_back * time.Hour)

	filtered_entries := []fs.DirEntry{}
	for _, entry := range entries {

		entry_date, err := get_entry_date(entry.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "error getting the entry date for %s", entry.Name())
		}

		if start_date.Before(entry_date) {
			filtered_entries = append(filtered_entries, entry)
		}
	}

	return filtered_entries, nil
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
