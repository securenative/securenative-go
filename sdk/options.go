package sdk

import "github.com/securenative/securenative-go/config"

func WithApiKey(apiKey string) config.SecureNativeOptions {
	options := config.DefaultSecureNativeOptions()
	options.ApiKey = apiKey
	return options
}
func WithOptions(options config.SecureNativeOptions) config.SecureNativeOptions {
	return options
}
func WithConfigFile(configPath string) config.SecureNativeOptions {
	configManager := config.NewConfigurationManager()
	options := configManager.LoadConfig(configPath)
	return options
}
