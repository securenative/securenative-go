package errors

type SecureNativeHttpError struct {
	Msg string
}

func (e *SecureNativeHttpError) Error() string {
	return e.Msg
}