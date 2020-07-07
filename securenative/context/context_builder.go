package context

import (
	"github.com/securenative/securenative-go/securenative/utils"
	"net/http"
)

const SecureNativeCookie = "_sn"

type ContextBuilderInterface interface {
	FromHttpRequest(request *http.Request) SecureNativeContext
}

type SecureNativeContextBuilder struct {
	Context *SecureNativeContext
}

func NewSecureNativeContextBuilder() *SecureNativeContextBuilder {
	return &SecureNativeContextBuilder{Context: &SecureNativeContext{
		ClientToken: "",
		Ip:          "",
		RemoteIp:    "",
		Headers:     nil,
		Url:         "",
		Method:      "",
		Body:        "",
	}}
}

func (c *SecureNativeContextBuilder) FromHttpRequest(request *http.Request) *SecureNativeContext {
	u := utils.Utils{}
	requestUtils := utils.RequestUtils{}
	cookie, err := request.Cookie(SecureNativeCookie)
	clientToken := ""
	if err == nil && cookie != nil {
		clientToken = cookie.Value
	}

	headers := parseHeaders(request)
	if u.IsNilOrEmpty(clientToken) {
		clientToken = requestUtils.GetSecureHeaderFromRequest(request)
	}

	return c.WithUrl(request.URL.String()).
		WithMethod(request.Method).
		WithHeaders(headers).
		WithClientToken(clientToken).
		WithIp(requestUtils.GetClientIpFromRequest(request)).
		WithRemoteIp(requestUtils.GetRemoteIpFromRequest(request)).
		WithBody("").Build()
}

func (c *SecureNativeContextBuilder) WithClientToken(clientToken string) *SecureNativeContextBuilder {
	c.Context.ClientToken = clientToken
	return c
}

func (c *SecureNativeContextBuilder) WithIp(ip string) *SecureNativeContextBuilder {
	c.Context.Ip = ip
	return c
}

func (c *SecureNativeContextBuilder) WithRemoteIp(remoteIp string) *SecureNativeContextBuilder {
	c.Context.RemoteIp = remoteIp
	return c
}

func (c *SecureNativeContextBuilder) WithHeaders(headers map[string][]string) *SecureNativeContextBuilder {
	c.Context.Headers = headers
	return c
}

func (c *SecureNativeContextBuilder) WithUrl(url string) *SecureNativeContextBuilder {
	c.Context.Url = url
	return c
}

func (c *SecureNativeContextBuilder) WithMethod(method string) *SecureNativeContextBuilder {
	c.Context.Method = method
	return c
}

func (c *SecureNativeContextBuilder) WithBody(body string) *SecureNativeContextBuilder {
	c.Context.Body = body
	return c
}

func (c *SecureNativeContextBuilder) Build() *SecureNativeContext {
	return c.Context
}

func parseHeaders(request *http.Request) map[string][]string {
	headers := map[string][]string{}
	for name, values := range request.Header {
		headers[name] = values
	}

	return headers
}
