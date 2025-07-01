package timeUtils

import (
	"time"
)

func converToUnix(isoTimestamp string) (int64, error) {
	t, err := time.Parse(time.RFC3339, isoTimestamp)
	if err != nil {
		panic(err)
	}
	return t.Unix(), nil
}
