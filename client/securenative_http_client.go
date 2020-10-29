package client

import (
	"bytes"
	"fmt"
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/logger"
	"github.com/securenative/securenative-go/utils"
	"net/http"
	"time"
)

const AuthorizationHeader = "Authorization"
const VersionHeader = "SN-Version"
const UserAgentHeader = "User-Agent"
const UserAgentHeaderValue = "SecureNative-go"
const ContentTypeHeader = "Content-Type"
const ContentTypeHeaderValue = "application/json"

type HttpClientInterface interface {
	Post(path string, body string) map[string]string
}

type SecureNativeHttpClient struct {
	Options config.SecureNativeOptions
}

func NewSecureNativeHttpClient(options config.SecureNativeOptions) *SecureNativeHttpClient {
	return &SecureNativeHttpClient{Options: options}
}

func (c *SecureNativeHttpClient) Post(path string, body []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.Options.ApiUrl, path)
	log := logger.GetLogger()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to build request; %s", err))
		return nil, err
	}

	for key, value := range c.GetHeaders() {
		req.Header.Add(key, value)
	}

	client := &http.Client{Timeout: time.Duration(c.Options.Timeout / 1000)}
	res, err := client.Do(req)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed to post request to %s; %s", c.Options.ApiUrl, err))
		return nil, err
	}

	return res, nil
}

func (c *SecureNativeHttpClient) GetHeaders() map[string]string {
	versionUtils := utils.VersionUtils{}
	return map[string]string{
		ContentTypeHeader:   ContentTypeHeaderValue,
		UserAgentHeader:     UserAgentHeaderValue,
		VersionHeader:       versionUtils.GetVersion(),
		AuthorizationHeader: c.Options.ApiKey,
	}
}