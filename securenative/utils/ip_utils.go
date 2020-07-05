package utils

import (
	"net"
	"regexp"
)

var ValidIpv4Pattern, _ = regexp.Compile("(([01]?\\d\\d?|2[0-4]\\d|25[0-5])\\.){3}([01]?\\d\\d?|2[0-4]\\d|25[0-5])")
var ValidIpv6Pattern, _ = regexp.Compile("([0-9a-f]{1,4}:){7}([0-9a-f]){1,4}")

type IpUtils struct{}

func NewIpUtils() *IpUtils {
	return &IpUtils{}
}

func (u *IpUtils) IsIpAddress(ip string) bool {
	if ValidIpv4Pattern.Match([]byte(ip)) {
		return true
	}
	if ValidIpv6Pattern.Match([]byte(ip)) {
		return true
	}
	return false
}

func (u *IpUtils) IsValidPublicIp(ip string) bool {
	parsedIp := net.ParseIP(ip)
	private := false
	if parsedIp == nil {
		return false
	} else {
		_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
		_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
		_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
		private = private24BitBlock.Contains(parsedIp) || private20BitBlock.Contains(parsedIp) || private16BitBlock.Contains(parsedIp)
	}

	if private || parsedIp.IsLoopback() || parsedIp.IsMulticast() || parsedIp.IsUnspecified() {
		return false
	}
	return true
}

func (u *IpUtils) IsLoopBack(ip string) bool {
	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		return false
	}
	return parsedIp.IsLoopback()
}
