package securenative

import (
	. "github.com/securenative/securenative-go/securenative/config"
	. "github.com/securenative/securenative-go/securenative/errors"
	. "github.com/securenative/securenative-go/securenative/models"
	. "github.com/securenative/securenative-go/securenative/utils"
)

type SDKInterface interface {
	Track(event SDKEvent)
	Verify(event SDKEvent)
	VerifyRequestPayload(event SDKEvent)
}

type SecureNative struct {
	options      SecureNativeOptions
	eventManager *EventManager
	apiManager   *ApiManager
	logger       *SdKLogger
}

var secureNative SecureNative

func NewSecureNative(options SecureNativeOptions) (*SecureNative, error) {
	utils := Utils{}
	if utils.IsNilOrEmpty(options.ApiKey) {
		return nil, &SecureNativeSDKError{Msg: "You must pass your SecureNative api key"}
	}

	secureNative := &SecureNative{}
	secureNative.options = options
	secureNative.eventManager = NewEventManager(options, nil)

	if len(options.ApiUrl) > 0 && options.ApiUrl != "" {
		secureNative.eventManager.StartEventPersist()
	}

	secureNative.apiManager = NewApiManager(secureNative.eventManager, options)
	secureNative.logger = InitLogger(options.LogLevel)

	return secureNative, nil
}

func (s *SecureNative) Init() error {
	panic("implement me")
}

func (s *SecureNative) InitSDKWithOptions(options SecureNativeOptions) error {
	panic("implement me")
}

func (s *SecureNative) InitSDKWithApiKey(apiKey string) error {
	panic("implement me")
}

func (s *SecureNative) Track(event SDKEvent) {
	panic("implement me")
}

func (s *SecureNative) Verify(event SDKEvent) {
	panic("implement me")
}

func (s *SecureNative) VerifyRequestPayload(event SDKEvent) {
	panic("implement me")
}

func (s *SecureNative) ConfigBuilder() *ConfigurationBuilder {
	return NewConfigurationBuilder()
}