package tests

import (
	"github.com/securenative/securenative-go/context"
	"github.com/securenative/securenative-go/sdk"
	"net/http"
	"net/url"
	"testing"
)

func TestCreateContextFromRequest(t *testing.T) {
	clientToken := "71532c1fad2c7f56118f7969e401f3cf080239140d208e7934e6a530818c37e544a0c2330a487bcc6fe4f662a57f265a3ed9f37871e80529128a5e4f2ca02db0fb975ded401398f698f19bb0cafd68a239c6caff99f6f105286ab695eaf3477365bdef524f5d70d9be1d1d474506b433aed05d7ed9a435eeca357de57817b37c638b6bb417ffb101eaf856987615a77a"
	headers := map[string][]string{
		"x-securenative": {clientToken},
	}
	request := &http.Request{
		Method: "Post",
		URL: &url.URL{
			Scheme: "https",
			Host:   "www.securenative.com",
			Path:   "/login",
		},
		Header:     headers,
		Host:       "www.securenative.com",
		RemoteAddr: "51.68.201.122",
		RequestURI: "/login/param1=value1&param2=value2",
	}

	c := sdk.FromHttpRequest(request)
	if c.ClientToken != clientToken {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", clientToken, c.ClientToken)
	}
	if c.Ip != "51.68.201.122" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "51.68.201.122", c.Ip)
	}
	if c.Method != "Post" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "Post", c.Method)
	}
	if c.RemoteIp != "51.68.201.122" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "51.68.201.122", c.RemoteIp)
	}
}

func TestCreateContextFromRequestWithCookie(t *testing.T) {
	clientToken := "71532c1fad2c7f56118f7969e401f3cf080239140d208e7934e6a530818c37e544a0c2330a487bcc6fe4f662a57f265a3ed9f37871e80529128a5e4f2ca02db0fb975ded401398f698f19bb0cafd68a239c6caff99f6f105286ab695eaf3477365bdef524f5d70d9be1d1d474506b433aed05d7ed9a435eeca357de57817b37c638b6bb417ffb101eaf856987615a77a"
	cookie := &http.Cookie{Name: "_sn", Value: clientToken}
	request := &http.Request{
		Method: "Post",
		URL: &url.URL{
			Scheme: "https",
			Host:   "www.securenative.com",
			Path:   "/login",
		},
		Header:     map[string][]string{},
		Host:       "www.securenative.com",
		RemoteAddr: "51.68.201.122",
		RequestURI: "/login/param1=value1&param2=value2",
	}
	request.AddCookie(cookie)

	c := sdk.FromHttpRequest(request)
	if c.ClientToken != clientToken {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", clientToken, c.ClientToken)
	}
	if c.Ip != "51.68.201.122" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "51.68.201.122", c.Ip)
	}
	if c.Method != "Post" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "Post", c.Method)
	}
	if c.RemoteIp != "51.68.201.122" {
		t.Errorf("Test Failed: expected to recieve: %s, got: %s", "51.68.201.122", c.RemoteIp)
	}
}

func TestCreateDefaultContextBuilder(t *testing.T) {
	c := context.SecureNativeContext{}

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
	c := context.SecureNativeContext{
		ClientToken: "SECRET_TOKEN",
		Ip:          "10.0.0.0",
		RemoteIp:    "10.0.0.1",
		Headers: map[string]string{
			"header1": "value",
		},
		Url:    "/some-url",
		Method: "GET",
		Body:   "{ \"name\": \"YOUR_NAME\" }",
	}

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
