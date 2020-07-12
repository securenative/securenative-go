package tests

import (
	"github.com/securenative/securenative-go/config"
	"github.com/securenative/securenative-go/enums"
	"github.com/securenative/securenative-go/sdk"
	"reflect"
	"testing"
)

func TestGetSdkInstanceWithoutInitThrows(t *testing.T) {
	sdk.Flush()

	_, err := sdk.GetInstance()
	if err == nil {
		t.Error("Test Failed: expected SecureNativeSDKIllegalStateError error to be thrown")
	}
}

func TestInitSdkWithoutApiKeyShouldThrow(t *testing.T) {
	sdk.Flush()

	_, err := sdk.InitSDK(config.SecureNativeOptions{})
	if err == nil {
		t.Error("Test Failed: expected SecureNativeSDKError error to be thrown")
	}
}

func TestInitDdkWithEmptyApiKeyShouldThrow(t *testing.T) {
	sdk.Flush()

	_, err := sdk.InitSDK(config.SecureNativeOptions{})
	if err == nil {
		t.Error("Test Failed: expected SecureNativeSDKError error to be thrown")
	}
}

func TestInitSdkWithApiKeyAndDefaults(t *testing.T) {
	sdk.Flush()

	apiKey := "SomeApiKey"
	s, err := sdk.InitSDK(sdk.WithApiKey(apiKey))

	if err != nil {
		t.Errorf("Test Failed: expected clean init; got error: %s", err)
	}

	options := s.GetSecureNativeOptions()
	if options.ApiKey != apiKey {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", apiKey, options.ApiKey)
	}
	if options.ApiUrl != "https://api.securenative.com/collector/api/v1" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "https://api.securenative.com/collector/api/v1", options.ApiUrl)
	}
	if options.Interval != 1000 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d", 1000, options.Interval)
	}
	if options.Timeout != 1500 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d", 1500, options.Timeout)
	}
	if options.MaxEvents != 1000 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d",1000, options.MaxEvents)
	}
	if options.AutoSend != true {
		t.Errorf("Test Failed: expected to reiecve: %t, got: %t", true, options.AutoSend)
	}
	if options.Disable != false {
		t.Errorf("Test Failed: expected to reiecve: %t, got: %t", false, options.Disable)
	}
	if options.LogLevel != "CRITICAL" {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", "CRITICAL", options.LogLevel)
	}
	if options.FailOverStrategy != enums.FailOverStrategy.FailOpen {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", enums.FailOverStrategy.FailOpen, options.FailOverStrategy)
	}
}

func TestInitSdkTwiceWillThrow(t *testing.T) {
	sdk.Flush()

	apiKey := "SomeApiKey"
	_, err := sdk.InitSDK(sdk.WithApiKey(apiKey))
	_, err = sdk.InitSDK(sdk.WithApiKey(apiKey))

	if err == nil {
		t.Error("Test Failed: expected SecureNativeSDKError error to be thrown")
	}
}

func TestInitSdkWithApiKeyAndGetInstance(t *testing.T) {
	sdk.Flush()

	apiKey := "SomeApiKey"
	s, err := sdk.InitSDK(sdk.WithApiKey(apiKey))

	if err != nil {
		t.Errorf("Test Failed: expected clean init; got error: %s", err)
	}

	instance, _ := sdk.GetInstance()
	if !reflect.DeepEqual(s, instance) {
		t.Errorf("Test Failed: expected both instances to be equal; sdk: %v; instance: %v", s, instance)
	}
}

func TestInitSdkWithBuilder(t *testing.T) {
	sdk.Flush()

	apiKey := "SomeApiKey"
	apiUrl := "SomeApiUrl"
	configBuilder := config.NewConfigurationBuilder()
	options := configBuilder.
		WithApiKey(apiKey).
		WithAutoSend(false).
		WithInterval(10).
		WithApiUrl(apiUrl).Build()

	s, err := sdk.InitSDK(sdk.WithOptions(options))
	if err != nil {
		t.Errorf("Test Failed: expected clean init; got error: %s", err)
	}

	o := s.GetSecureNativeOptions()
	if o.ApiKey != apiKey {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", apiKey, o.ApiKey)
	}
	if o.ApiUrl != apiUrl {
		t.Errorf("Test Failed: expected to reiecve: %s, got: %s", apiUrl, o.ApiUrl)
	}
	if o.Interval != 10 {
		t.Errorf("Test Failed: expected to reiecve: %d, got: %d", 1000, o.Interval)
	}
	if o.AutoSend != false {
		t.Errorf("Test Failed: expected to reiecve: %t, got: %t", true, o.AutoSend)
	}
}
