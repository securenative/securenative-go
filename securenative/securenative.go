package securenative

import (
	. "github.com/securenative/securenative-go/securenative/config"
	. "github.com/securenative/securenative-go/securenative/models"
)

type SDKInterface interface {
	Track(event SDKEvent)
	Verify(event SDKEvent)
	VerifyRequestPayload(event SDKEvent)
}

type SecureNative struct {
	// TODO implement me
}

func NewSecureNative() (*SecureNative, error) {
	panic("implement me")
}

func InitSDK() (*SecureNative, error) {
	panic("implement me")
}

func InitSDKWithOptions(options SecureNativeOptions) (*SecureNative, error) {
	panic("implement me")
}

func InitSDKWithApiKey(apiKey string) (*SecureNative, error) {
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
