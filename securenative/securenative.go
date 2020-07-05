package securenative

import (
	. "github.com/securenative/securenative-go/securenative/config"
	. "github.com/securenative/securenative-go/securenative/context"
	. "github.com/securenative/securenative-go/securenative/errors"
	. "github.com/securenative/securenative-go/securenative/models"
	. "github.com/securenative/securenative-go/securenative/utils"
	"io/ioutil"
	. "net/http"
	"runtime"
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

var secureNative *SecureNative

func newSecureNative(options SecureNativeOptions) (*SecureNative, error) {
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

	runtime.SetFinalizer(secureNative, exit)

	return secureNative, nil
}

func InitSDK() (*SecureNative, error) {
	if secureNative != nil {
		return secureNative, &SecureNativeSDKError{Msg: "This SDK was already initialized"}
	}

	configManager := NewConfigurationManager()
	options := configManager.LoadConfig()
	sn, err := newSecureNative(options)

	if err != nil {
		return nil, err
	}

	secureNative = sn
	return sn, nil
}

func InitSDKWithOptions(options SecureNativeOptions) (*SecureNative, error) {
	if secureNative != nil {
		return secureNative, &SecureNativeSDKError{Msg: "This SDK was already initialized"}
	}

	sn, err := newSecureNative(options)

	if err != nil {
		return nil, err
	}

	secureNative = sn
	return sn, nil
}

func InitSDKWithApiKey(apiKey string) (*SecureNative, error) {
	if secureNative != nil {
		return secureNative, &SecureNativeSDKError{Msg: "This SDK was already initialized"}
	}

	utils := Utils{}
	if utils.IsNilOrEmpty(apiKey) {
		return nil,  &SecureNativeConfigError{Msg: "You must pass your SecureNative api key"}
	}

	builder := NewConfigurationBuilder()
	options := builder.WithApiKey(apiKey).Build()
	sn, err := newSecureNative(options)

	if err != nil {
		return nil, err
	}

	secureNative = sn
	return sn, nil
}

func (s *SecureNative) Track(event EventOptions) {
	s.apiManager.Track(event)
}

func (s *SecureNative) Verify(event EventOptions) VerifyResult {
	return s.apiManager.Verify(event)
}

func (s *SecureNative) VerifyRequestPayload(request *Request) bool {
	signatureUtils := NewSignatureUtils()

	requestSignature := request.Header.Get(SignatureHeader)
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		s.logger.Debug("Could not read request body")
		body = []byte("")
	}

	return signatureUtils.IsValidSignature(s.options.ApiKey, string(body), requestSignature)
}

func (s *SecureNative) ConfigBuilder() *ConfigurationBuilder {
	return NewConfigurationBuilder()
}

func (s *SecureNative) ContextBuilder() *ContextBuilder {
	return NewContextBuilder()
}

func (s *SecureNative) SecureNativeOptions() SecureNativeOptions {
	return s.options
}

func GetInstance() *SecureNative {
	return secureNative
}

func (s *SecureNative) ReleaseSDK() {
	if secureNative != nil {
		s.eventManager.StopEventPersist()
		secureNative = nil
	}
}