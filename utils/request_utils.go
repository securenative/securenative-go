package utils

import (
	"github.com/securenative/securenative-go/config"
	"net/http"
	"strings"
)

var ipHeaders = []string{"x-forwarded-for", "x-client-ip", "x-real-ip", "x-forwarded", "x-cluster-client-ip", "forwarded-for", "forwarded", "via"}
var ipUtils = NewIpUtils()

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
			values := request.Header[header]
			for _, ip := range values {
				if len(ip) > 0 || ip != "" {
					if strings.Contains(ip, ",") {
						ips := strings.Split(ip, ",")
						extracted := getValidIp(ips)
						if extracted != "" {
							return extracted
						}
					} else {
						candidate := ipUtils.NormalizeIp(ip)
						if ipUtils.IsValidPublicIp(candidate) {
							return candidate
						}
						if !ipUtils.IsLoopBack(candidate) {
							return candidate
						}
					}
				}
			}
		}
	}

	for _, header := range ipHeaders {
		if ips, ok := request.Header[header]; ok {
			for _, ip := range ips {
				if strings.Contains(ip, ",") {
					ips := strings.Split(ip, ",")
					extracted := getValidIp(ips)
					if extracted != "" {
						return extracted
					}
				} else {
					candidate := ipUtils.NormalizeIp(ip)
					if ipUtils.IsValidPublicIp(candidate) {
						return candidate
					}
					if !ipUtils.IsLoopBack(candidate) {
						return candidate
					}
				}
			}
		}
	}

	return request.RemoteAddr
}

func (u *RequestUtils) GetRemoteIpFromRequest(request *http.Request) string {
	return request.RemoteAddr
}

func getValidIp(ips []string) string {
	for _, extracted := range ips {
		ip := ipUtils.NormalizeIp(extracted)
		if ipUtils.IsValidPublicIp(ip) {
			return ip
		}
	}
	for _, extracted := range ips {
		ip := ipUtils.NormalizeIp(extracted)
		if !ipUtils.IsLoopBack(ip) {
			return ip
		}
	}
	return ""
}
