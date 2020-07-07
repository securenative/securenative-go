package tests

import (
	"github.com/securenative/securenative-go/securenative/context"
	"testing"
)

func TestCreateContextFromRequest(t *testing.T) {
	// TODO implement me
}

func TestCreateContextFromRequestWithCookie(t *testing.T) {
	// TODO implement me
}

func TestCreateDefaultContextBuilder(t *testing.T) {
	c := context.NewSecureNativeContextBuilder().Build()

	if c.Url != "" {
		t.Errorf("Test Failed: expected to recieve an empty field, got: %s", c.Url)
	}
	if c.ClientToken != "" {
		t.Errorf("Test Failed: expected to recieve an empty field, got: %s", c.ClientToken)
	}
	if c.Ip != "" {
		t.Errorf("Test Failed: expected to recieve an empty field, got: %s", c.Ip)
	}
	if c.Body != "" {
		t.Errorf("Test Failed: expected to recieve an empty field, got: %s", c.Body)
	}
	if c.Method != "" {
		t.Errorf("Test Failed: expected to recieve an empty field, got: %s", c.Method)
	}
	if c.RemoteIp != "" {
		t.Errorf("Test Failed: expected to recievean empty field, got: %s", c.RemoteIp)
	}
}

func TestCreateCustomContextWithContextBuilder(t *testing.T) {
	c := context.NewSecureNativeContextBuilder().
		WithUrl("/some-url").
		WithClientToken("SECRET_TOKEN").
		WithIp("10.0.0.0").
		WithBody("{ \"name\": \"YOUR_NAME\" }").
		WithMethod("GET").
		WithRemoteIp("10.0.0.1").
		WithHeaders(map[string][]string{
			"header1": {"value1", "value2"},
		}).Build()

	if c.Url != "/some-url" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "/some-url", c.Url)
	}
	if c.ClientToken != "SECRET_TOKEN" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "SECRET_TOKEN", c.ClientToken)
	}
	if c.Ip != "10.0.0.0" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "10.0.0.0", c.Ip)
	}
	if c.Body != "{ \"name\": \"YOUR_NAME\" }" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "{ \"name\": \"YOUR_NAME\" }", c.Body)
	}
	if c.Method != "GET" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "GET", c.Method)
	}
	if c.RemoteIp != "10.0.0.1" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "10.0.0.1", c.RemoteIp)
	}
}
