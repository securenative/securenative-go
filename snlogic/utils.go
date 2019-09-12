package snlogic

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/mergermarket/go-pkcs7"
	"io"
)

const AES_KEY_SIZE = 32

// Encrypt encrypts plain text string into cipher text string
func Encrypt(unencrypted string, cipherKey string) (string, error) {
	SnLog(fmt.Sprintf("Starting encrypt %s ", unencrypted))
	key := []byte(TrimKey(cipherKey))
	plainText := []byte(unencrypted)
	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)
	if err != nil {
		SnLog(fmt.Sprintf("encryption padding  failed: %s ", unencrypted))
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		SnLog(fmt.Sprintf("encryption wrong block size %s ", unencrypted))
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		SnLog(fmt.Sprintf("encryption failed %s ", unencrypted))
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)
	SnLog(fmt.Sprintf("finished encryption %s ", unencrypted))
	return fmt.Sprintf("%x", cipherText), nil
}

// Decrypt decrypts cipher text string into plain text string
func Decrypt(encrypted string, cipherKey string) (string, error) {
	key := []byte(TrimKey(cipherKey))
	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		SnLog(fmt.Sprintf("decryption failed"))
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		SnLog(fmt.Sprintf("decryption failed cipherText too short"))
		return "", errors.New("decryption failed cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		SnLog(fmt.Sprintf("decryption failed cipherText is not a multiple of the block size"))
		return "", errors.New("decryption failed cipherText is not a multiple of the block size")

	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	SnLog(fmt.Sprintf("decryption completed"))
	return fmt.Sprintf("%s", cipherText), nil
}

func TrimKey(cipherKey string) string {
	return cipherKey[:AES_KEY_SIZE]
}
