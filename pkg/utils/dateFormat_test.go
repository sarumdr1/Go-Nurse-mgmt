package utils

import (
	"testing"
)

func TestFormatDate(t *testing.T) {
	date := "2023-09-07 16:55:41.79958 +0545 +0545"
	expectedDate := "Sep 09,2023"
	formattedDate := FormatDate(date)

	if formattedDate != expectedDate {
		t.Errorf("Expected Date %s doesn't match with %s", expectedDate, formattedDate)
	}
}
