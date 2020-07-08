package errors

type SecureNativeSDKIllegalStateError struct {
	Msg string
}

func (e *SecureNativeSDKIllegalStateError) Error() string {
	return e.Msg
}
