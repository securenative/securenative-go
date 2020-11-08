package tests

import (
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/utils"
	"net/http"
	"testing"
)

func TestProxyHeadersExtractionFromRequestIpv4 (t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header:     map[string][]string{"CF-Connecting-IP": {"203.0.113.1"}},
	}

	options := config.DefaultSecureNativeOptions()
	options.ProxyHeaders = []string{"CF-Connecting-IP"}

	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "203.0.113.1" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "203.0.113.1")
	}
}

func TestProxyHeadersExtractionFromRequestIpv6 (t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header:     map[string][]string{"CF-Connecting-IP": {"f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d"}},
	}

	options := config.DefaultSecureNativeOptions()
	options.ProxyHeaders = []string{"CF-Connecting-IP"}

	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "f71f:5bf9:25ff:1883:a8c4:eeff:7b80:aa2d")
	}
}

func TestProxyHeadersExtractionFromRequestMultipleIps (t *testing.T) {
	requestUtils := utils.NewRequestUtils()
	request := &http.Request{
		Header:     map[string][]string{"CF-Connecting-IP": {"141.246.115.116, 203.0.113.1, 12.34.56.3"}},
	}

	options := config.DefaultSecureNativeOptions()
	options.ProxyHeaders = []string{"CF-Connecting-IP"}

	clientIp := requestUtils.GetClientIpFromRequest(request, &options)

	if clientIp != "141.246.115.116" {
		t.Errorf("Test Failed: extracted ip is: %s, expected: %s", clientIp, "141.246.115.116")
	}
}