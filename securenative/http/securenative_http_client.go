package http

import (
	"bytes"
	"fmt"
	. "github.com/securenative/securenative-go/securenative"
	. "github.com/securenative/securenative-go/securenative/config"
	. "github.com/securenative/securenative-go/securenative/utils"
	. "net/http"
)

const AuthorizationHeader = "Authorization"
const VersionHeader = "SN-Version"
const UserAgentHeader = "User-Agent"
const UserAgentHeaderValue = "SecureNative-python"
const ContentTypeHeader = "Content-Type"
const ContentTypeHeaderValue = "application/json"

type HttpClient interface {
	Post(path string, body string) map[string]string
}

type SecureNativeHttpClient struct {
	Options SecureNativeOptions
}

func NewSecureNativeHttpClient(options SecureNativeOptions) *SecureNativeHttpClient {
	return &SecureNativeHttpClient{Options: options}
}

func (c *SecureNativeHttpClient) Post(path string, body string) *Response {
	url := fmt.Sprintf("%s/%s", c.Options.ApiUrl, path)
	logger := GetLogger()

	var jsonBody = []byte(body)
	req, err := NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		logger.Debug(fmt.Sprintf("Failed to build request; %s", err))
		return nil
	}

	for key, value := range c.GetHeaders() {
		req.Header.Add(key, value)
	}

	client := &Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Debug(fmt.Sprintf("Failed to post request to %s; %s", c.Options.ApiUrl, err))
		return nil
	}

	return res
}

func (c *SecureNativeHttpClient) GetHeaders() map[string]string {
	versionUtils := VersionUtils{}
	return map[string]string{
		ContentTypeHeader:   ContentTypeHeaderValue,
		UserAgentHeader:     UserAgentHeaderValue,
		VersionHeader:       versionUtils.GetVersion(),
		AuthorizationHeader: c.Options.ApiKey,
	}
}