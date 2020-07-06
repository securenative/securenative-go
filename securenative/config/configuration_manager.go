package config

import (
	"fmt"
	"github.com/securenative/securenative-go/securenative/utils"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

const DefaultConfigFile = "securenative.yml"
const CustomConfigFileEnvName = "SECURENATIVE_COMFIG_FILE"

type ConfigurationManagerInterface interface {
	LoadConfig() SecureNativeOptions
}

type ConfigurationManager struct{}

func NewConfigurationManager() *ConfigurationManager {
	return &ConfigurationManager{}
}

func (c *ConfigurationManager) LoadConfig() SecureNativeOptions {
	configurationBuilder := NewConfigurationBuilder()
	options := configurationBuilder.DefaultSecureNativeOptions()

	resourcePath := DefaultConfigFile
	if len(os.Getenv(CustomConfigFileEnvName)) > 0 {
		resourcePath = DefaultConfigFile
	}

	properties := c.readResourceFile(resourcePath)

	return configurationBuilder.
		WithApiKey(c.getStringEnvOrDefault(properties, "SECURENATIVE_API_KEY", options.ApiKey)).
		WithApiUrl(c.getStringEnvOrDefault(properties, "SECURENATIVE_API_URL", options.ApiUrl)).
		WithInterval(c.getIntEnvOrDefault(properties, "SECURENATIVE_INTERVAL", options.Interval)).
		WithMaxEvents(c.getIntEnvOrDefault(properties, "SECURENATIVE_MAX_EVENTS", options.MaxEvents)).
		WithTimeout(c.getIntEnvOrDefault(properties, "SECURENATIVE_TIMEOUT", options.Timeout)).
		WithAutoSend(c.getBoolEnvOrDefault(properties, "SECURENATIVE_AUTO_SEND", options.AutoSend)).
		WithDisable(c.getBoolEnvOrDefault(properties, "SECURENATIVE_DISABLE", options.Disable)).
		WithLogLevel(c.getStringEnvOrDefault(properties, "SECURENATIVE_LOG_LEVEL", options.LogLevel)).
		WithLogLevel(c.getStringEnvOrDefault(properties, "SECURENATIVE_FAILOVER_STRATEGY", options.FailOverStrategy)).Build()
}

func (c *ConfigurationManager) readResourceFile(path string) map[string]string {
	logger := utils.GetLogger()

	file, err := os.Open(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Could not read file %s; %s", path, err))
	}

	var cfg map[string]string
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Could decode file %s; %s", path, err))
	}

	if file != nil {
		defer file.Close()
	}

	return cfg
}

func (c *ConfigurationManager) getStringEnvOrDefault(properties map[string]string, key string, defaultKey string) string {
	if len(os.Getenv(key)) > 0 {
		return os.Getenv(key)
	}

	if len(properties[key]) > 0 {
		return properties[key]
	}

	return defaultKey
}

func (c *ConfigurationManager) getIntEnvOrDefault(properties map[string]string, key string, defaultKey int32) int32 {
	if len(os.Getenv(key)) > 0 {
		data, err := strconv.Atoi(os.Getenv(key))
		if err != nil {
			return defaultKey
		}
		return int32(data)
	}

	if len(properties[key]) > 0 {
		data, err := strconv.Atoi(properties[key])
		if err != nil {
			return defaultKey
		}
		return int32(data)
	}

	return defaultKey
}

func (c *ConfigurationManager) getBoolEnvOrDefault(properties map[string]string, key string, defaultKey bool) bool {
	if len(os.Getenv(key)) > 0 {
		data, err := strconv.ParseBool(os.Getenv(key))
		if err != nil {
			return defaultKey
		}
		return data
	}

	if len(properties[key]) > 0 {
		data, err := strconv.ParseBool(properties[key])
		if err != nil {
			return defaultKey
		}
		return data
	}

	return defaultKey
}
