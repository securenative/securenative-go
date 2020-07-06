package tests

import (
	"github.com/securenative/securenative-go/securenative/utils"
	"testing"
)

func TestToTimestamp(t *testing.T) {
	dateUtils := utils.NewDateUtils()
	Iso8601Date := "2020-05-20T15:07:13Z"
	result := dateUtils.ToTimestamp(Iso8601Date)

	expected := "2020-05-20T15:07:13Z"

	if result != expected {
		t.Errorf("Test Failed: %s inputted, %s expected; %s received", Iso8601Date, expected, result)
	}
}
