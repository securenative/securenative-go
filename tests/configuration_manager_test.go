package tests

import (
	"github.com/securenative/securenative-go/securenative/config"
	"testing"
)

func TestParseConfigFileCorrectly(t *testing.T) {
	configManager := config.NewConfigurationManager()
	options := configManager.LoadConfig()

	if options.ApiKey != "Some Random Api Key" {
		t.Error("Test Failed: configuration file was not read correctly")
	}
}

func TestIgnoreUnknownConfigInPropertiesFile(t *testing.T) {
	// TODO implement me
}

func TestHandleInvalidConfigFile(t *testing.T) {
	// TODO implement me
}

func TestIgnoreInvalidConfigFileEntries(t *testing.T) {
	// TODO implement me
}

func TestLoadDefaultConfig(t *testing.T) {
	// TODO implement me
}

func TestGetConfigFromEnvVariables(t *testing.T) {
	// TODO implement me
}

func TestDefaultValuesForInvalidEnumConfigProps(t *testing.T) {
	// TODO implement me
}
