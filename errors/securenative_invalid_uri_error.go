package errors

type SecureNativeInvalidUriError struct {
	Msg string
}

func (e *SecureNativeInvalidUriError) Error() string {
	return e.Msg
}
