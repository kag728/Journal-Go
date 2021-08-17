package entry_utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func get_entry_date(filename string) (time.Time, error) {

	split_date := strings.Split(filename, "_")
	month_index := 1

	if len(split_date) == Max_entry_name_sections {
		month_index = 2
	}

	month, err := get_month_from_string(split_date[month_index])
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "error converting month to time.Month")
	}
	day, err := strconv.Atoi(split_date[month_index+1])
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "error converting day to int")
	}
	year, err := strconv.Atoi(split_date[month_index+2])
	if err != nil {
		return time.Time{}, errors.Wrapf(err, "error converting year to int")
	}

	date := time.Date(year, month, day, 1, 1, 1, 1, time.Local)

	return date, nil
}

func get_days_back(day string) int {

	switch day {
	case "Monday":
		return 1
	case "Tuesday":
		return 2
	case "Wednesday":
		return 3
	case "Thursday":
		return 4
	case "Friday":
		return 5
	case "Saturday":
		return 6
	case "Sunday":
		return 0
	}

	return 0
}

func get_month_from_string(month string) (time.Month, error) {
	switch month {
	case "January":
		return time.January, nil
	case "February":
		return time.February, nil
	case "March":
		return time.March, nil
	case "April":
		return time.April, nil
	case "May":
		return time.May, nil
	case "June":
		return time.June, nil
	case "July":
		return time.July, nil
	case "August":
		return time.August, nil
	case "September":
		return time.September, nil
	case "October":
		return time.October, nil
	case "November":
		return time.November, nil
	case "December":
		return time.December, nil
	}
	return time.January, fmt.Errorf("could not convert %s to a month", month)
}
