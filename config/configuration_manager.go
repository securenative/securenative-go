package config

import (
	"fmt"
	"github.com/securenative/securenative-go/logger"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
	"strings"
)

const DefaultConfigFile = "securenative.yml"
const CustomConfigFileEnvName = "SECURENATIVE_CONFIG_FILE"

type ConfigurationManagerInterface interface {
	LoadConfig() SecureNativeOptions
}

type ConfigurationManager struct{}

func NewConfigurationManager() *ConfigurationManager {
	return &ConfigurationManager{}
}

func (c *ConfigurationManager) LoadConfig(configPath string) SecureNativeOptions {
	options := DefaultSecureNativeOptions()

	resourcePath := DefaultConfigFile
	if len(configPath) > 1 && configPath != "" {
		resourcePath = configPath
	}
	if len(os.Getenv(CustomConfigFileEnvName)) > 0 {
		resourcePath = os.Getenv(CustomConfigFileEnvName)
	}

	properties := c.readResourceFile(resourcePath)

	return SecureNativeOptions{
		ApiKey:           c.getStringEnvOrDefault(properties, "SECURENATIVE_API_KEY", options.ApiKey),
		ApiUrl:           c.getStringEnvOrDefault(properties, "SECURENATIVE_API_URL", options.ApiUrl),
		Interval:         c.getIntEnvOrDefault(properties, "SECURENATIVE_INTERVAL", options.Interval),
		MaxEvents:        c.getIntEnvOrDefault(properties, "SECURENATIVE_MAX_EVENTS", options.MaxEvents),
		Timeout:          c.getIntEnvOrDefault(properties, "SECURENATIVE_TIMEOUT", options.Timeout),
		AutoSend:         c.getBoolEnvOrDefault(properties, "SECURENATIVE_AUTO_SEND", options.AutoSend),
		Disable:          c.getBoolEnvOrDefault(properties, "SECURENATIVE_DISABLE", options.Disable),
		LogLevel:         c.getStringEnvOrDefault(properties, "SECURENATIVE_LOG_LEVEL", options.LogLevel),
		FailOverStrategy: c.getStringEnvOrDefault(properties, "SECURENATIVE_FAILOVER_STRATEGY", options.FailOverStrategy),
		ProxyHeaders:     c.getSliceEnvOrDefault(properties, "SECURENATIVE_PROXY_HEADERS", options.ProxyHeaders),
		PiiHeaders:       c.getSliceEnvOrDefault(properties, "SECURENATIVE_PII_HEADERS", options.PiiHeaders),
		PiiRegexPattern:  c.getStringEnvOrDefault(properties, "SECURENATIVE_PII_REGEX_PATTERN", options.PiiRegexPattern),
	}
}

func (c *ConfigurationManager) readResourceFile(path string) map[string]string {
	log := logger.GetLogger()

	file, err := os.Open(path)
	if err != nil {
		log.Debug(fmt.Sprintf("Could not read file %s; %s", path, err))
	}

	var cfg map[string]string
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Debug(fmt.Sprintf("Could decode file %s; %s", path, err))
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

func (c *ConfigurationManager) getSliceEnvOrDefault(properties map[string]string, key string, defaultKey []string) []string {
	if len(os.Getenv(key)) > 0 {
		return parseStringSlice(os.Getenv(key))
	}

	if len(properties[key]) > 0 {
		return parseStringSlice(properties[key])
	}

	return defaultKey
}

func parseStringSlice(str string) []string {
	s1 := strings.Replace(str, "[", "", -1)
	s2 := strings.Replace(s1, "]", "", -1)
	return strings.Split(s2, ",")
}
