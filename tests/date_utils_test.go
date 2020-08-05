package tests

import (
	"github.com/securenative/securenative-go/utils"
	"testing"
	"time"
)

func TestToTimestamp(t *testing.T) {
	dateUtils := utils.NewDateUtils()
	Iso8601Date := "2020-05-20T15:07:13Z"
	ts, _ := time.Parse("2006-01-02T15:04:05.999Z" , Iso8601Date)
	result := dateUtils.ToTimestamp(ts)

	expected := "2020-05-20T15:07:13Z"

	if result != expected {
		t.Errorf("Test Failed: %s inputted, %s expected; %s received", Iso8601Date, expected, result)
	}
}
