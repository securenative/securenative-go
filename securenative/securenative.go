package securenative

import (
	"github.com/securenative/securenative-go/securenative/config"
	"github.com/securenative/securenative-go/securenative/errors"
	"github.com/securenative/securenative-go/securenative/events"
	"github.com/securenative/securenative-go/securenative/models"
	"github.com/securenative/securenative-go/securenative/utils"
	"io/ioutil"
	"net/http"
)

type SDKInterface interface {
	Track(event models.SDKEvent)
	Verify(event models.SDKEvent)
	VerifyRequestPayload(event models.SDKEvent)
}

type SecureNative struct {
	options        config.SecureNativeOptions
	eventManager   *events.EventManager
	apiManager     *events.ApiManager
	logger         *utils.SdKLogger
}

var secureNative *SecureNative

func newSecureNative(options config.SecureNativeOptions) (*SecureNative, error) {
	u := utils.Utils{}
	if u.IsNilOrEmpty(options.ApiKey) {
		return nil, &errors.SecureNativeSDKError{Msg: "You must pass your SecureNative api key"}
	}

	secureNative := &SecureNative{}
	secureNative.options = options
	secureNative.eventManager = events.NewEventManager(options, nil)

	if len(options.ApiUrl) > 0 && options.ApiUrl != "" {
		secureNative.eventManager.StartEventPersist()
	}

	secureNative.apiManager = events.NewApiManager(secureNative.eventManager, options)
	secureNative.logger = utils.InitLogger(options.LogLevel)

	return secureNative, nil
}

func InitSDK() (*SecureNative, error) {
	if secureNative != nil {
		return secureNative, &errors.SecureNativeSDKError{Msg: "This SDK was already initialized"}
	}

	configManager := config.NewConfigurationManager()
	options := configManager.LoadConfig()
	sn, err := newSecureNative(options)

	if err != nil {
		return nil, err
	}

	secureNative = sn
	return sn, nil
}

func InitSDKWithOptions(options config.SecureNativeOptions) (*SecureNative, error) {
	if secureNative != nil {
		return secureNative, &errors.SecureNativeSDKError{Msg: "This SDK was already initialized"}
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
		return secureNative, &errors.SecureNativeSDKError{Msg: "This SDK was already initialized"}
	}

	u := utils.Utils{}
	if u.IsNilOrEmpty(apiKey) {
		return nil, &errors.SecureNativeConfigError{Msg: "You must pass your SecureNative api key"}
	}

	configBuilder := config.NewConfigurationBuilder()
	options := configBuilder.WithApiKey(apiKey).Build()
	sn, err := newSecureNative(options)

	if err != nil {
		return nil, err
	}

	secureNative = sn
	return sn, nil
}

func (s *SecureNative) Track(event models.EventOptions) {
	s.apiManager.Track(event)
}

func (s *SecureNative) Verify(event models.EventOptions) *models.VerifyResult {
	return s.apiManager.Verify(event)
}

func (s *SecureNative) VerifyRequestPayload(request *http.Request) bool {
	signatureUtils := utils.NewSignatureUtils()

	requestSignature := request.Header.Get(utils.SignatureHeader)
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		s.logger.Debug("Could not read request body")
		body = []byte("")
	}

	return signatureUtils.IsValidSignature(s.options.ApiKey, string(body), requestSignature)
}

func (s *SecureNative) GetEventOptionsBuilder(eventType string) *events.EventOptionsBuilder {
	return events.NewEventOptionsBuilder(eventType)
}

func (s *SecureNative) GetSecureNativeOptions() config.SecureNativeOptions {
	return s.options
}

func GetInstance() (*SecureNative, error) {
	if secureNative == nil {
		return secureNative, &errors.SecureNativeSDKIllegalStateError{Msg: "SDK was not initialized"}
	}
	return secureNative, nil
}

func (s *SecureNative) Stop() {
	if secureNative != nil {
		s.eventManager.StopEventPersist()
		secureNative = nil
	}
}

func Flush() {
	secureNative = nil
}
