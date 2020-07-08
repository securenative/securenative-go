package enums

type ApiRouteEnum struct {
	Track  string
	Verify string
}

var ApiRoute = ApiRouteEnum{
	Track:  "track",
	Verify: "verify",
}
