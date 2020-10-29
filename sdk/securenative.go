package sdk

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/errors"
	"github.com/securenative/securenative-go/events"
	"github.com/securenative/securenative-go/models"
	"github.com/securenative/securenative-go/utils"
)

type SDKInterface interface {
	Track(event models.SDKEvent) error
	Verify(event models.SDKEvent) (*models.VerifyResult, error)
	VerifyRequestPayload(request *http.Request) bool
}

type SecureNative struct {
	options      config.SecureNativeOptions
	eventManager *events.EventManager
	apiManager   *events.ApiManager
	logger       *utils.SdKLogger
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		secureNative.eventManager.StopEventPersist()
	}()

	return secureNative, nil
}

func InitSDK(options config.SecureNativeOptions) (*SecureNative, error) {
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

func (s *SecureNative) Track(event models.EventOptions) error {
	return s.apiManager.Track(event)
}

func (s *SecureNative) Verify(event models.EventOptions) (*models.VerifyResult, error) {
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

func (s *SecureNative) GetSecureNativeOptions() config.SecureNativeOptions {
	return s.options
}

func GetInstance() (*SecureNative, error) {
	if secureNative == nil {
		return secureNative, &errors.SecureNativeSDKIllegalStateError{Msg: "SDK was not initialized"}
	}
	return secureNative, nil
}

func Flush() {
	secureNative = nil
}
