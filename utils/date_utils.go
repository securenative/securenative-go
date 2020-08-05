package utils

import (
	"time"
)

type DateUtils struct{}

func NewDateUtils() *DateUtils {
	return &DateUtils{}
}

func (u *DateUtils) ToTimestamp(timestamp time.Time) string {
	return timestamp.Format("2006-01-02T15:04:05.999Z")
}
