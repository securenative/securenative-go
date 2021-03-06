package tests

import (
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/enums"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var f, _ = filepath.Abs("securenative.yml")
var configPath = strings.Replace(f, "/tests", "", -1)

func createConfigFile() {
	f, _ := os.Create(configPath)
	defer f.Close()
	options := config.SecureNativeOptions{
		ApiKey: "Some Random Api Key",
	}

	out, _ := yaml.Marshal(options)
	_, _ = f.Write(out)
}

func TestParseConfigFileCorrectly(t *testing.T) {
	createConfigFile()
	configManager := config.NewConfigurationManager()
	options := configManager.LoadConfig(configPath)

	if options.ApiKey != "Some Random Api Key" {
		t.Error("Test Failed: configuration file was not read correctly")
	}
}

func TestLoadDefaultConfig(t *testing.T) {
	configManager := config.NewConfigurationManager()
	options := configManager.LoadConfig("")

	if options.ApiKey != "" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "", options.ApiKey)
	}
	if options.ApiUrl != "https://api.securenative.com/collector/api/v1" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "https://api.securenative.com/collector/api/v1", options.ApiUrl)
	}
	if options.Interval != 1000 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d", 1000, options.Interval)
	}
	if options.Timeout != 1500 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d", 1500, options.Timeout)
	}
	if options.MaxEvents != 1000 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d",1000, options.MaxEvents)
	}
	if options.AutoSend != true {
		t.Errorf("Test Failed: expected to reiecve: %t, got: %t", true, options.AutoSend)
	}
	if options.Disable != false {
		t.Errorf("Test Failed: expected to reiecve: %t, got: %t", false, options.Disable)
	}
	if options.LogLevel != "CRITICAL" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "CRITICAL", options.LogLevel)
	}
	if options.FailOverStrategy != enums.FailOverStrategy.FailOpen {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", enums.FailOverStrategy.FailOpen, options.FailOverStrategy)
	}

	if len(options.ProxyHeaders) !=0 {
		t.Errorf("Test Failed: expected lenght of proxy headers to be: %d, got: %d", 0, len(options.ProxyHeaders))
	}
}

func TestGetConfigFromEnvVariables(t *testing.T) {
	configManager := config.NewConfigurationManager()

	_ = os.Setenv("SECURENATIVE_API_KEY", "SOME_ENV_API_KEY")
	_ = os.Setenv("SECURENATIVE_API_URL", "SOME_API_URL")
	_ = os.Setenv("SECURENATIVE_INTERVAL", "6000")
	_ = os.Setenv("SECURENATIVE_MAX_EVENTS", "700")
	_ = os.Setenv("SECURENATIVE_TIMEOUT", "1700")
	_ = os.Setenv("SECURENATIVE_AUTO_SEND", "false")
	_ = os.Setenv("SECURENATIVE_DISABLE", "true")
	_ = os.Setenv("SECURENATIVE_LOG_LEVEL", "DEBUG")
	_ = os.Setenv("SECURENATIVE_FAILOVER_STRATEGY", "fail-closed")
	_ = os.Setenv("SECURENATIVE_PROXY_HEADERS", "CF-Connecting-Ip,Some-Random-Ip")
	_ = os.Setenv("SECURENATIVE_PII_HEADERS", "authentication,apiKey")
	_ = os.Setenv("SECURENATIVE_PII_REGEX_PATTERN", "/http_auth_/i")

	options := configManager.LoadConfig("")
	expectedProxyHeaders := []string{"CF-Connecting-Ip", "Some-Random-Ip"}
	os.Clearenv()

	if options.ApiKey != "SOME_ENV_API_KEY" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "SOME_ENV_API_KEY", options.ApiKey)
	}
	if options.ApiUrl != "SOME_API_URL" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "SOME_API_URL", options.ApiUrl)
	}
	if options.Interval != 6000 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d", 6000, options.Interval)
	}
	if options.Timeout != 1700 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d", 1700, options.Timeout)
	}
	if options.MaxEvents != 700 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d",700, options.MaxEvents)
	}
	if options.AutoSend != false {
		t.Errorf("Test Failed: expected to reiecve: %t, got: %t", false, options.AutoSend)
	}
	if options.Disable != true {
		t.Errorf("Test Failed: expected to reiecve: %t, got: %t", true, options.Disable)
	}
	if options.LogLevel != "DEBUG" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "DEBUG", options.LogLevel)
	}
	if options.PiiRegexPattern != "/http_auth_/i" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "/http_auth_/i", options.PiiRegexPattern)
	}
	if options.FailOverStrategy != enums.FailOverStrategy.FailClose {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", enums.FailOverStrategy.FailClose, options.FailOverStrategy)
	}
	if len(options.ProxyHeaders) != len(expectedProxyHeaders) {
		t.Errorf("Test Failed: expected length of proxy headers to be: %d, got: %d", len([]string{"CF-Connecting-Ip", "Some-Random-Ip"}),len(options.ProxyHeaders))
	}
	if options.ProxyHeaders[0] != "CF-Connecting-Ip" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "CF-Connecting-Ip", options.ProxyHeaders[0])
	}
	if options.ProxyHeaders[1] != "Some-Random-Ip" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "Some-Random-Ip", options.ProxyHeaders[1])
	}
	if options.PiiHeaders[0] != "authentication" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "authentication", options.PiiHeaders[0])
	}
	if options.PiiHeaders[1] != "apiKey" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "apiKey", options.PiiHeaders[1])
	}
}