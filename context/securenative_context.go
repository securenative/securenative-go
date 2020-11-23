package context

import (
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/utils"
	"net/http"
	"strings"
)

const SecureNativeCookie = "_sn"
var piiHeaders = []string{"authorization", "access_token", "apikey", "password", "passwd", "secret", "api_key"}

type SecureNativeContext struct {
	ClientToken string
	Ip          string
	RemoteIp    string
	Headers     map[string]string
	Url         string
	Method      string
	Body        string
}

func FromHttpRequest(request *http.Request, options *config.SecureNativeOptions) *SecureNativeContext {
	u := utils.Utils{}
	requestUtils := utils.RequestUtils{}
	cookie, err := request.Cookie(SecureNativeCookie)
	clientToken := ""
	if err == nil && cookie != nil {
		clientToken = cookie.Value
	}

	headers := ParseHeaders(request)
	if u.IsNilOrEmpty(clientToken) {
		clientToken = requestUtils.GetSecureHeaderFromRequest(request)
	}

	return &SecureNativeContext{
		ClientToken: clientToken,
		Ip:          requestUtils.GetClientIpFromRequest(request, options),
		RemoteIp:    requestUtils.GetRemoteIpFromRequest(request),
		Headers:     headers,
		Url:         request.URL.String(),
		Method:      request.Method,
		Body:        "",
	}
}

func ParseHeaders(request *http.Request) map[string]string {
	headers := map[string]string{}
	for name, values := range request.Header {
		if !contains(piiHeaders, name) && !contains(piiHeaders, strings.ToUpper(name)) {
			headers[name] = values[0]
		}
	}

	return headers
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}