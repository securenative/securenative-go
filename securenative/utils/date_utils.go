package utils

import (
	"time"
)

type DateUtils struct{}

func NewDateUtils() *DateUtils {
	return &DateUtils{}
}

func (u *DateUtils) ToTimestamp(timestamp string) string {
	ts, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return time.Now().Format(time.RFC3339)
	}

	return ts.String()
}
