package utils

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (u *Utils) IsNilOrEmpty(str string) bool {
	if len(str) == 0 || str == "" {
		return true
	}
	return false
}
