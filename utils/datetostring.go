package utils

import (
	"time"
)

func DateToString(t time.Time) (string, error) {
	// const inputLayout = "2006-01-02T15:04:05.000000-07:00"
	const outputLayout = "02-01-2006"
	// t, err := time.Parse(inputLayout, dateStr)
	// if err != nil {
	// 	return "", err
	// }
	// const outputLayout = "02-01-2006"
	// formattedDate := t.Format(outputLayout)

	return t.Format(outputLayout), nil
}
