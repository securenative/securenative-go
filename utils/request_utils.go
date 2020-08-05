package utils

import "net/http"

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
