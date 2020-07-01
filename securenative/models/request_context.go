package models

type RequestContext struct {
	Cid      string
	Vid      string
	Fp       string
	Ip       string
	RemoteIp string
	Headers  map[string][]string
	Url      string
	Method   string
}

