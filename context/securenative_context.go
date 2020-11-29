package context

import (
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/utils"
	"net/http"
)

const SecureNativeCookie = "_sn"

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

	headers := requestUtils.GetHeadersFromRequest(request, options)
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