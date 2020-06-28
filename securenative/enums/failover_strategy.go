package enums

type FailOverStrategy string

const (
	FAIL_OPEN   FailOverStrategy = "fail-open"
	FAIL_CLOSED FailOverStrategy = "fail-closed"
)
