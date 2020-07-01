package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
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

func (u *EncryptionUtils) Decrypt(encrypted string, cipherKey string) (ClientTokenObject, error) {
	key := []byte(TrimKey(cipherKey))
	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		return ClientTokenObject{}, err
	}

	if len(cipherText) < aes.BlockSize {
		return ClientTokenObject{}, errors.New("decryption failed; cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		return ClientTokenObject{}, errors.New("decryption failed cipherText is not a multiple of the block size")

	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	var decrypted ClientTokenObject
	err = json.Unmarshal(cipherText, &decrypted)
	if err != nil {
		return ClientTokenObject{}, fmt.Errorf("failed to marshal encrypted text; %s", err)
	}

	return decrypted, nil
}

func (u *EncryptionUtils) Encrypt(text string, cipherKey string) (string, error) {
	key := []byte(TrimKey(cipherKey))
	plainText := []byte(text)
	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf("could not pad plain text; %s", plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf("wrong block size was given for %s", plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)
	return fmt.Sprintf("%x", cipherText), nil
}

func TrimKey(cipherKey string) string {
	return cipherKey[:AesKeySize]
}