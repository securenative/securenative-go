package sdk

import "github.com/securenative/securenative-go/config"

func WithApiKey(apiKey string) config.SecureNativeOptions {
	configBuilder := config.NewConfigurationBuilder()
	options := configBuilder.WithApiKey(apiKey).Build()
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
