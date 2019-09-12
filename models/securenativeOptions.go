package models

type SecureNativeOptions struct {
	ApiUrl           string `envDefault:"https://api.securenative.com/collector/api/v1"`
	Timeout          int64  `envDefault:1500`
	Interval         int    `envDefault:1000`
	MaxEvents        int    `envDefault:1000`
	AutoSend         bool   `envDefault:true`
	IsSdkEnabled     bool   `envDefault:true`
	IsLoggingEnabled bool   `envDefault:false`
}
