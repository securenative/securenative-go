package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type SignatureUtils struct{}

const SignatureHeader = "x-securenative"

func NewSignatureUtils() *SignatureUtils {
	return &SignatureUtils{}
}

func (u *SignatureUtils) IsValidSignature(apiKey string, payload string, headerSignature string) bool {

	key := []byte(apiKey)
	body := []byte(payload)

	h := hmac.New(sha256.New, key)
	h.Write(body)
	signature := hex.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(headerSignature))
}
