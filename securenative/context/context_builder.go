package context

import (
	. "github.com/securenative/securenative-go/securenative/utils"
	. "net/http"
)

const SecureNativeCookie = "_sn"

type ContextBuilderInterface interface {
	FromHttpRequest(request *Request) SecureNativeContext
}

type ContextBuilder struct {
	Context *SecureNativeContext
}

func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{Context: &SecureNativeContext{
		ClientToken: "",
		Ip:          "",
		RemoteIp:    "",
		Headers:     nil,
		Url:         "",
		Method:      "",
		Body:        "",
	}}
}

func (c *ContextBuilder) FromHttpRequest(request *Request) *SecureNativeContext {
	utils := Utils{}
	requestUtils := RequestUtils{}
	cookie, err := request.Cookie(SecureNativeCookie)
	clientToken := ""
	if err == nil {
		clientToken = cookie.Value
	}

	headers := parseHeaders(request)
	if utils.IsNilOrEmpty(clientToken) {
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

func (c *ContextBuilder) WithClientToken(clientToken string) *ContextBuilder {
	c.Context.ClientToken = clientToken
	return c
}

func (c *ContextBuilder) WithIp(ip string) *ContextBuilder {
	c.Context.Ip = ip
	return c
}

func (c *ContextBuilder) WithRemoteIp(remoteIp string) *ContextBuilder {
	c.Context.RemoteIp = remoteIp
	return c
}

func (c *ContextBuilder) WithHeaders(headers map[string][]string) *ContextBuilder {
	c.Context.Headers = headers
	return c
}

func (c *ContextBuilder) WithUrl(url string) *ContextBuilder {
	c.Context.Url = url
	return c
}

func (c *ContextBuilder) WithMethod(method string) *ContextBuilder {
	c.Context.Method = method
	return c
}

func (c *ContextBuilder) WithBody(body string) *ContextBuilder {
	c.Context.Body = body
	return c
}

func (c *ContextBuilder) Build() *SecureNativeContext {
	return c.Context
}

func parseHeaders(request *Request) map[string][]string {
	headers := map[string][]string{}
	for name, values := range request.Header {
		headers[name] = values
	}

	return headers
}
