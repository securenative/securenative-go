package config

var DEFAULT_CONFIG_FILE = "securenative.ini"
var CUSTOM_CONFIG_FILE_ENV_NAME = "SECURENATIVE_COMFIG_FILE"

type ConfigurationManagerInterface interface {
	ConfigBuilder() *ConfigurationBuilder
	LoadConfig() SecureNativeOptions
}

type ConfigurationManager struct {
	// TODO implement me
}

func NewConfigurationManager() *ConfigurationManager {
	panic("implement me")
}

func (c *ConfigurationManager) ConfigBuilder() *ConfigurationBuilder {
	panic("implement me")
}

func (c *ConfigurationManager) LoadConfig() SecureNativeOptions {
	panic("implement me")
}