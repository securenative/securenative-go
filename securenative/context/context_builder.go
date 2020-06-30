package context

import "net/http"

type ContextBuilder struct {
	// TODO implement me
}

func NewContextBuilder() *ContextBuilder {
	panic("implement me")
}

func (c *ContextBuilder) DefaultContextBuilder() (*ContextBuilder, error) {
	panic("implement me")
}

func (c *ContextBuilder) FromHttpRequest(request http.Request) (*ContextBuilder, error) {
	panic("implement me")
}