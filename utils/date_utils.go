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
		return time.Now().Format("2006-01-02T15:04:05Z")
	}

	return ts.Format("2006-01-02T15:04:05Z")
}
