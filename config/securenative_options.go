package config

import "github.com/securenative/securenative-go/enums"

type SecureNativeOptions struct {
	ApiKey           string   `yaml:"SECURENATIVE_API_KEY"`
	ApiUrl           string   `yaml:"SECURENATIVE_API_URL"`
	Interval         int32    `yaml:"SECURENATIVE_INTERVAL"`
	MaxEvents        int32    `yaml:"SECURENATIVE_MAX_EVENTS"`
	Timeout          int32    `yaml:"SECURENATIVE_TIMEOUT"`
	AutoSend         bool     `yaml:"SECURENATIVE_AUTO_SEND"`
	Disable          bool     `yaml:"SECURENATIVE_DISABLE"`
	LogLevel         string   `yaml:"SECURENATIVE_LOG_LEVEL"`
	FailOverStrategy string   `yaml:"SECURENATIVE_FAILOVER_STRATEGY"`
	ProxyHeaders     []string `yaml:"SECURENATIVE_PROXY_HEADERS"`
	PiiHeaders       []string `yaml:"SECURENATIVE_PII_HEADERS"`
	PiiRegexPattern  string   `yaml:"SECURENATIVE_PII_REGEX_PATTERN"`
}

func DefaultSecureNativeOptions() SecureNativeOptions {
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
		ProxyHeaders:     []string{},
		PiiHeaders:       []string{},
		PiiRegexPattern:  "",
	}
}
