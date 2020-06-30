package context

import "net/http"

type ContextBuilderInterface interface {
	FromHttpRequest(request http.Request) SecureNativeContext
}

type ContextBuilder struct {
	Context SecureNativeContext
}

func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{Context: SecureNativeContext{
		ClientToken: "",
		Ip:          "",
		RemoteIp:    "",
		Headers:     nil,
		Url:         "",
		Method:      "",
		Body:        "",
	}}
}

func (c *ContextBuilder) FromHttpRequest(request http.Request) SecureNativeContext {
	panic("implement me")
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

func (c *ContextBuilder) WithHeaders(headers map[string]string) *ContextBuilder {
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

func (c *ContextBuilder) Build() SecureNativeContext {
	return c.Context
}