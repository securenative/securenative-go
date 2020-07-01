package http

import . "github.com/securenative/securenative-go/securenative/config"

const AUTHORIZATION_HEADER = "Authorization"
const VERSION_HEADER = "SN-Version"
const USER_AGENT_HEADER = "User-Agent"
const USER_AGENT_HEADER_VALUE = "SecureNative-python"
const CONTENT_TYPE_HEADER = "Content-Type"
const CONTENT_TYPE_HEADER_VALUE = "application/json"

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