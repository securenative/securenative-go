package utils

type EncryptionUtils struct {}

type ClientTokenObject struct {
	Cid string
	Vid string
	Fp  string
}

func (u EncryptionUtils) Decrypt(secret string, apiKey string) ClientTokenObject {
	// TODO implement me
	return ClientTokenObject{
		Cid: "",
		Vid: "",
		Fp:  "",
	}
}

func (u EncryptionUtils) Encrypt() bool {
	// TODO implement me
	return false
}
