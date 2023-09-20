package utils

import (
	"time"
)

func FormatDate(date string) string {
	format := "2006-01-02 15:04:05.99999 -0700 -0700"
	parsedDate, err := time.Parse(format, date)

	if err != nil {
		return ""
	}
	formattedDate := parsedDate.Format("Jan 01,2006")
	return formattedDate

}
