package tests

import (
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/context"
	"github.com/securenative/securenative-go/utils"
	"net/http"
	"testing"
)

func TestProxyHeadersExtractionFromRequestIpv4(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"CF-Connecting-IP": {"203.0.113.1"}},
	}

	options := config.DefaultSecureNativeOptions()
	options.ProxyHeaders = []string{"CF-Connecting-IP"}

	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "203.0.113.1" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "203.0.113.1")
	}
}

func TestProxyHeadersExtractionFromRequestIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"CF-Connecting-IP": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	options.ProxyHeaders = []string{"CF-Connecting-IP"}

	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestProxyHeadersExtractionFromRequestMultipleIps(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"CF-Connecting-IP": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	options.ProxyHeaders = []string{"CF-Connecting-IP"}

	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingXForwardedForHeaderIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-forwarded-for": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingXForwardedForHeaderMultipleIp(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-forwarded-for": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingXClientIpHeaderIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-client-ip": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingXClientIpHeaderMultipleIp(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-client-ip": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingXForwardedHeaderIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-forwarded": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingXForwardedHeaderMultipleIp(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-forwarded": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingXRealIpHeaderIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-real-ip": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingXRealIpHeaderMultipleIp(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-real-ip": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingXxClusterClientIpHeaderIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-cluster-client-ip": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingXxClusterClientIpHeaderMultipleIp(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"x-cluster-client-ip": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingForwardedForHeaderIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"forwarded-for": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingForwardedForHeaderMultipleIp(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"forwarded-for": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingForwardedHeaderIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"forwarded": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingForwardedHeaderMultipleIp(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"forwarded": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingViaHeaderIpv6(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"via": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingViaHeaderMultipleIp(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{"via": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestIpExtractionUsingPriorityWithXForwardedFor(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{
			"x-forwarded-for": {"203.0.113.1"},
			"x-real-ip":       {"198.51.100.101"},
			"x-client-ip":     {"198.51.100.102"},
		},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "203.0.113.1" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestIpExtractionUsingPriorityWithoutXForwardedFor(t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header: map[string][]string{
			"x-real-ip":   {"198.51.100.101"},
			"x-client-ip": {"203.0.113.1, 141.246.115.116, 12.34.56.3"},
		},
	}

	options := config.DefaultSecureNativeOptions()
	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "203.0.113.1" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}

func TestPiiDataExtractionFromHeaders(t *testing.T) {
	headers := map[string][]string{
		"Host":            {"net.example.com"},
		"User-Agent":      {"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5 (.NET CLR 3.5.30729)"},
		"Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"Accept-Language": {"en-us,en;q=0.5"},
		"Accept-Encoding": {"gzip,deflate"},
		"Accept-Charset":  {"ISO-8859-1,utf-8;q=0.7,*;q=0.7"},
		"Keep-Alive":      {"300"},
		"Connection":      {"keep-alive"},
		"Cookie":          {"PHPSESSID=r2t5uvjq435r4q7ib3vtdjq120"},
		"Pragma":          {"no-cache"},
		"Cache-Control":   {"no-cache"},
		"authorization":   {"ylSkZIjbdWybfs4fUQe9BqP0LH5Z"},
		"access_token":    {"ylSkZIjbdWybfs4fUQe9BqP0LH5Z"},
		"apikey":          {"ylSkZIjbdWybfs4fUQe9BqP0LH5Z"},
		"password":        {"ylSkZIjbdWybfs4fUQe9BqP0LH5Z"},
		"passwd":          {"ylSkZIjbdWybfs4fUQe9BqP0LH5Z"},
		"secret":          {"ylSkZIjbdWybfs4fUQe9BqP0LH5Z"},
		"api_key":         {"ylSkZIjbdWybfs4fUQe9BqP0LH5Z"},
	}
	request := &http.Request{
		Header: headers,
	}

	h := context.ParseHeaders(request)
	if Contains(h, "authorization") != false {
		t.Errorf("Test Failed: extracted header: %s, not to be present", "authorization")
	}

	if Contains(h, "access_token") != false {
		t.Errorf("Test Failed: extracted header: %s, not to be present", "access_token")
	}

	if Contains(h, "apikey") != false {
		t.Errorf("Test Failed: extracted header: %s, not to be present", "apikey")
	}

	if Contains(h, "password") != false {
		t.Errorf("Test Failed: extracted header: %s, not to be present", "password")
	}

	if Contains(h, "secret") != false {
		t.Errorf("Test Failed: extracted header: %s, not to be present", "secret")
	}

	if Contains(h, "api_key") != false {
		t.Errorf("Test Failed: extracted header: %s, not to be present", "api_key")
	}
}

func Contains(s map[string]string, e string) bool {
	for key := range s {
		if key == e {
			return true
		}
	}
	return false
}