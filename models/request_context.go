package models

type RequestContext struct {
	Cid         string            `json:"cid"`
	Vid         string            `json:"vid"`
	Fp          string            `json:"fp"`
	Ip          string            `json:"ip"`
	RemoteIp    string            `json:"remoteIp"`
	Headers     map[string]string `json:"headers"`
	Url         string            `json:"url"`
	Method      string            `json:"method"`
	ClientToken string            `json:"clientToken"`
}
