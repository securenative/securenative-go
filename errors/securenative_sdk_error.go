package errors

type SecureNativeSDKError struct {
	Msg string
}

func (e *SecureNativeSDKError) Error() string {
	return e.Msg
}
