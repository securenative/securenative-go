package tests

import (
	"github.com/securenative/securenative-go/securenative/utils"
	"testing"
)

func TestIsIpAddressValidIpv4(t *testing.T) {
	ipUtils := utils.NewIpUtils()
	validIpv4 := "172.16.254.1"

	result := ipUtils.IsIpAddress(validIpv4)

	if result != true {
		t.Errorf("Test Failed: %s inputted, %t expected; %t received", validIpv4, true, result)
	}
}

func TestIsIpAddressValidIpv6(t *testing.T) {
	ipUtils := utils.NewIpUtils()
	validIpv6 := "2001:db8:1234:0000:0000:0000:0000:0000"

	result := ipUtils.IsIpAddress(validIpv6)

	if result != true {
		t.Errorf("Test Failed: %s inputted, %t expected; %t received", validIpv6, true, result)
	}
}

func TestIsIpAddressInvalidIpv4(t *testing.T) {
	ipUtils := utils.NewIpUtils()
	invalidIpv4 := "172.16.2541"

	result := ipUtils.IsIpAddress(invalidIpv4)

	if result != false {
		t.Errorf("Test Failed: %s inputted, %t expected; %t received", invalidIpv4, false, result)
	}
}

func TestIsIpAddressInvalidIpv6(t *testing.T) {
	ipUtils := utils.NewIpUtils()
	invalidIpv6 := "2001:db8:1234:0000"

	result := ipUtils.IsIpAddress(invalidIpv6)

	if result != false {
		t.Errorf("Test Failed: %s inputted, %t expected; %t received", invalidIpv6, false, result)
	}
}

func TestIsValidPublicIp(t *testing.T) {
	ipUtils := utils.NewIpUtils()
	ip := "64.71.222.37"

	result := ipUtils.IsValidPublicIp(ip)

	if result != true {
		t.Errorf("Test Failed: %s inputted, %t expected; %t received", ip, true, result)
	}
}

func TestIsNotValidPublicIp(t *testing.T) {
	ipUtils := utils.NewIpUtils()
	ip := "10.0.0.0"

	result := ipUtils.IsValidPublicIp(ip)

	if result != false {
		t.Errorf("Test Failed: %s inputted, %t expected; %t received", ip, false, result)
	}
}

func TestIsValidLoopbackIp(t *testing.T) {
	ipUtils := utils.NewIpUtils()
	ip := "127.0.0.1"

	result := ipUtils.IsLoopBack(ip)

	if result != true {
		t.Errorf("Test Failed: %s inputted, %t expected; %t received", ip, true, result)
	}
}