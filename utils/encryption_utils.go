package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"pault.ag/go/pkcs7"
)

const AesKeySize = 32

type EncryptionUtils struct{}

func NewEncryptionUtils() *EncryptionUtils {
	return &EncryptionUtils{}
}

type ClientTokenObject struct {
	Cid string
	Vid string
	Fp  string
}

func (u *EncryptionUtils) Decrypt(encrypted string, cipherKey string) ClientTokenObject {
	if len(encrypted) >= AesKeySize && len(cipherKey) >= AesKeySize {
		key := []byte(TrimKey(cipherKey))
		cipherText, _ := hex.DecodeString(encrypted)

		block, err := aes.NewCipher(key)
		if err != nil {
			return ClientTokenObject{}
		}

		if len(cipherText) < aes.BlockSize {
			return ClientTokenObject{}
		}
		iv := cipherText[:aes.BlockSize]
		cipherText = cipherText[aes.BlockSize:]
		if len(cipherText)%aes.BlockSize != 0 {
			return ClientTokenObject{}

		}

		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(cipherText, cipherText)

		cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
		var decrypted ClientTokenObject
		err = json.Unmarshal(cipherText, &decrypted)
		if err != nil {
			return ClientTokenObject{}
		}

		return decrypted
	}
	return ClientTokenObject{}
}

func (u *EncryptionUtils) Encrypt(text string, cipherKey string) string {
	key := []byte(TrimKey(cipherKey))
	plainText := []byte(text)
	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)
	if err != nil {
		return ""
	}
	if len(plainText)%aes.BlockSize != 0 {
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)
	return fmt.Sprintf("%x", cipherText)
}

func TrimKey(cipherKey string) string {
	return cipherKey[:AesKeySize]
}