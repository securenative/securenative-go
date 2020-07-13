package context

type SecureNativeContext struct {
	ClientToken string
	Ip          string
	RemoteIp    string
	Headers     map[string]string
	Url         string
	Method      string
	Body        string
}
