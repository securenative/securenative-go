package utils

import "net/http"

const SecureNativeHeader = "x-securenative"

type RequestUtils struct{}

func NewRequestUtils() *RequestUtils {
	return &RequestUtils{}
}

func (u *RequestUtils) GetSecureHeaderFromRequest(request *http.Request) string {
	return request.Header.Get(SecureNativeHeader)
}

func (u *RequestUtils) GetClientIpFromRequest(request *http.Request) string {
	ip := request.Header.Get("X-Forwarded-For")
	if len(ip) == 0 || ip == "" {
		return request.RemoteAddr
	}
	return ip
}

func (u *RequestUtils) GetRemoteIpFromRequest(request *http.Request) string {
	return request.RemoteAddr
}
