package http

import . "github.com/securenative/securenative-go/securenative/config"

var AUTHORIZATION_HEADER = "Authorization"
var VERSION_HEADER = "SN-Version"
var USER_AGENT_HEADER = "User-Agent"
var USER_AGENT_HEADER_VALUE = "SecureNative-python"
var CONTENT_TYPE_HEADER = "Content-Type"
var CONTENT_TYPE_HEADER_VALUE = "application/json"

type HttpClient interface {
	Post(path string, body string) map[string]string
}

type SecureNativeHttpClient struct {
	// TODO implement me
}

func NewSecureNativeHttpClient(options SecureNativeOptions) *SecureNativeHttpClient {
	panic("implement me")
}

func (c *SecureNativeHttpClient) Post(path string, body string) map[string]string {
	panic("implement me")
}