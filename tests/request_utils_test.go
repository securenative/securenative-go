package tests

import (
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/utils"
	"net/http"
	"testing"
)

func TestProxyHeadersExtractionFromRequest (t *testing.T) {
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