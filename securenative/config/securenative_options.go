package config

type SecureNativeOptions struct {
	ApiKey           string
	ApiUrl           string
	Interval         int32
	MaxEvents        int32
	Timeout          int32
	AutoSend         bool
	Disable          bool
	LogLevel         string
	FailOverStrategy string
}
