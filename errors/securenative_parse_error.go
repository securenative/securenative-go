package errors

type SecureNativeParseError struct {
	Msg string
}

func (e *SecureNativeParseError) Error() string {
	return e.Msg
}
