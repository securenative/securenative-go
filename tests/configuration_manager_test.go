package tests

import (
	"github.com/securenative/securenative-go/securenative/config"
	"github.com/securenative/securenative-go/securenative/enums"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createConfigFile() {
	filename, _ := filepath.Abs("securenative.yml")
	filename = strings.Replace(filename, "/tests", "", -1)
	f, _ := os.Create(filename)
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
	options := configManager.LoadConfig()

	if options.ApiKey != "Some Random Api Key" {
		t.Error("Test Failed: configuration file was not read correctly")
	}
}

func TestLoadDefaultConfig(t *testing.T) {
	configManager := config.NewConfigurationManager()
	options := configManager.LoadConfig()

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

	options := configManager.LoadConfig()
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
	if options.FailOverStrategy != enums.FailOverStrategy.FailClose {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", enums.FailOverStrategy.FailClose, options.FailOverStrategy)
	}
}