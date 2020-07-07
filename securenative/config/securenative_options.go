package config

type SecureNativeOptions struct {
	ApiKey           string `yaml:"SECURENATIVE_API_KEY"`
	ApiUrl           string `yaml:"SECURENATIVE_API_URL"`
	Interval         int32  `yaml:"SECURENATIVE_INTERVAL"`
	MaxEvents        int32  `yaml:"SECURENATIVE_MAX_EVENTS"`
	Timeout          int32  `yaml:"SECURENATIVE_TIMEOUT"`
	AutoSend         bool   `yaml:"SECURENATIVE_AUTO_SEND"`
	Disable          bool   `yaml:"SECURENATIVE_DISABLE"`
	LogLevel         string `yaml:"SECURENATIVE_LOG_LEVEL"`
	FailOverStrategy string `yaml:"SECURENATIVE_FAILOVER_STRATEGY"`
}
