package models

import . "go/types"

type RequestContext struct {
	Cid      string
	Vid      string
	Fp       string
	Ip       string
	RemoteIp string
	Headers  Slice
	Url      string
	Method   string
}

