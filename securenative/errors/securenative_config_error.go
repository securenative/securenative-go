package errors

type SecureNativeConfigError struct {
	Msg string
}

func (e *SecureNativeConfigError) Error() string {
	return e.Msg
}
