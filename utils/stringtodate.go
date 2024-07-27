package utils

import (
	"fmt"
	"time"
)

func StringToDate(effectiveDate string) (time.Time, error) {
	if effectiveDate == "" {
		return time.Time{}, fmt.Errorf("effectiveDate string is empty")
	}
	layout := "02-01-2006" // Updated layout to match the format "DD-MM-YYYY"
	parsedTime, err := time.Parse(layout, effectiveDate)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date: %w", err)
	}
	return parsedTime, nil
}
