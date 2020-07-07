package tests

import (
	"github.com/securenative/securenative-go/securenative/config"
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
