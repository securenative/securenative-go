package config

import "github.com/securenative/securenative-go/securenative/enums"

type ConfigurationBuilder struct {
	Options SecureNativeOptions
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	options := SecureNativeOptions{
		ApiKey:           "",
		ApiUrl:           "https://api.securenative.com/collector/api/v1",
		Interval:         1000,
		MaxEvents:        1000,
		Timeout:          1500,
		AutoSend:         true,
		Disable:          false,
		LogLevel:         "CRITICAL",
		FailOverStrategy: enums.FailOverStrategy.FailOpen,
	}

	return &ConfigurationBuilder{Options: options}
}

func (c *ConfigurationBuilder) DefaultSecureNativeOptions() SecureNativeOptions {
	return SecureNativeOptions{
		ApiKey:           "",
		ApiUrl:           "https://api.securenative.com/collector/api/v1",
		Interval:         1000,
		MaxEvents:        1000,
		Timeout:          1500,
		AutoSend:         true,
		Disable:          false,
		LogLevel:         "CRITICAL",
		FailOverStrategy: enums.FailOverStrategy.FailOpen,
	}
}

func (c *ConfigurationBuilder) WithApiKey(apiKey string) *ConfigurationBuilder {
	c.Options.ApiUrl = apiKey
	return c
}

func (c *ConfigurationBuilder) WithApiUrl(apiUrl string) *ConfigurationBuilder {
	c.Options.ApiUrl = apiUrl
	return c
}

func (c *ConfigurationBuilder) WithInterval(interval int32) *ConfigurationBuilder {
	c.Options.Interval = interval
	return c
}

func (c *ConfigurationBuilder) WithMaxEvents(maxEvents int32) *ConfigurationBuilder {
	c.Options.MaxEvents = maxEvents
	return c
}

func (c *ConfigurationBuilder) WithTimeout(timeout int32) *ConfigurationBuilder {
	c.Options.Timeout = timeout
	return c
}

func (c *ConfigurationBuilder) WithAutoSend(autoSend bool) *ConfigurationBuilder {
	c.Options.AutoSend = autoSend
	return c
}

func (c *ConfigurationBuilder) WithDisable(disable bool) *ConfigurationBuilder {
	c.Options.Disable = disable
	return c
}

func (c *ConfigurationBuilder) WithLogLevel(logLevel string) *ConfigurationBuilder {
	c.Options.LogLevel = logLevel
	return c
}

func (c *ConfigurationBuilder) WithFailOverStrategy(failOverStrategy string) *ConfigurationBuilder {
	c.Options.FailOverStrategy = failOverStrategy
	return c
}

func (c *ConfigurationBuilder) Build() SecureNativeOptions {
	return c.Options
}
