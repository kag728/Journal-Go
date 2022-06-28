package entry_utils

import (
	"fmt"
	"io/fs"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (

	// MaxPrefixLength Number of digits before entry number (0000, 0001 is 4)
	MaxPrefixLength = 4

	// MaxEntryNameSections Number of underscore-separated sections in an entry file name (4 when weekday is not included)
	MaxEntryNameSections = 5
	MinEntryNameSections = 4

	// FileDir The directory containing the entries
	FileDir string = "entries"
)

// FilterEntries Filter list of entries so each on begins with a number and ends with a number.
// Definitely not a great implementation but we'll see how it does
func FilterEntries(entries []fs.DirEntry) []fs.DirEntry {

	var filteredEntries []fs.DirEntry
	for _, entry := range entries {

		entryNameSplit := strings.Split(entry.Name(), "_")

		if (len(entryNameSplit)) > MaxEntryNameSections || (len(entryNameSplit)) < MinEntryNameSections {
			continue
		}

		_, err := strconv.Atoi(entryNameSplit[0])
		if err != nil {
			continue
		}

		_, err = strconv.Atoi(entryNameSplit[len(entryNameSplit)-1])
		if err != nil {
			continue
		}

		filteredEntries = append(filteredEntries, entry)
	}

	return filteredEntries
}

func FilterEntriesForWeek(entries []fs.DirEntry) ([]fs.DirEntry, error) {

	entries = FilterEntries(entries)

	now := time.Now()
	today := now.Weekday()

	daysBack := getDaysBack(fmt.Sprint(today))
	hoursBack := time.Duration(daysBack * 24)

	startDate := now.Add(-hoursBack * time.Hour)

	var filteredEntries []fs.DirEntry
	for _, entry := range entries {

		entryDate, err := getEntryDate(entry.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "error getting the entry date for %s", entry.Name())
		}

		if startDate.Before(entryDate) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return filteredEntries, nil
}

func getPrefixDigits(n int) int {
	digits := 0
	for i := n; i >= 1; i = i / 10 {
		digits++
	}
	return digits
}

func FillPrefix(n int) (string, error) {

	if n == 0 {
		return "0000", nil
	}

	digits := getPrefixDigits(n)
	if digits > MaxPrefixLength {
		return "", errors.Wrapf(errors.New("Prefix error"), "Entry prefix too high. Max possible prefix: %d", MaxPrefixLength)
	}

	output := ""
	for i := 0; i < MaxPrefixLength-digits; i++ {
		output += "0"
	}

	return fmt.Sprintf("%s%d", output, n), nil
}
