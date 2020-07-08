package enums

type FailOverStrategyEnum struct {
	FailOpen  string
	FailClose string
}

var FailOverStrategy = FailOverStrategyEnum{
	FailOpen:  "fail-open",
	FailClose: "fail-closed",
}
