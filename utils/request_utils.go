package utils

import (
	"github.com/securenative/securenative-go/config"
	"net/http"
)

type RequestUtils struct{}

func NewRequestUtils() *RequestUtils {
	return &RequestUtils{}
}

func (u *RequestUtils) GetSecureHeaderFromRequest(request *http.Request) string {
	header := request.Header["x-securenative"]
	if len(header) >= 1 {
		return header[0]
	} else {
		return ""
	}
}

func (u *RequestUtils) GetClientIpFromRequest(request *http.Request, options *config.SecureNativeOptions) string {
	if options != nil && len(options.ProxyHeaders) > 0 {
		for _, header := range options.ProxyHeaders {
			ip := request.Header.Get(header)
			if len(ip) > 0 || ip != "" {
				return ip
			}
		}
	}
	
	ip := request.Header.Get("X-Forwarded-For")
	if len(ip) == 0 || ip == "" {
		return request.RemoteAddr
	}
	return ip
}

func (u *RequestUtils) GetRemoteIpFromRequest(request *http.Request) string {
	return request.RemoteAddr
}
