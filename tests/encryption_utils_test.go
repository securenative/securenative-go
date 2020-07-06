package tests

import (
	"github.com/securenative/securenative-go/securenative/utils"
	"testing"
)

var SecretKey = "B00C42DAD33EAC6F6572DA756EA4915349C0A4F6"
var Payload = "{\"cid\":\"198a41ff-a10f-4cda-a2f3-a9ca80c0703b\",\"vi\":\"148a42ff-b40f-4cda-a2f3-a8ca80c0703b\",\"fp\":\"6d8cabd95987f8318b1fe01593d5c2a5.24700f9f1986800ab4fcc880530dd0ed\"}"
var Cid = "198a41ff-a10f-4cda-a2f3-a9ca80c0703b"
var Fp = "6d8cabd95987f8318b1fe01593d5c2a5.24700f9f1986800ab4fcc880530dd0ed"

func TestEncrypt(t *testing.T) {
	encryptionUtils := utils.NewEncryptionUtils()
	expected := "885d1b8fc2b2ba3af9009d3c6bce307de8b6564edbe835795ac706ed82d4e4f61e79cfdebdb01268f56f23763d3590f1c0b36cf81c04b7e38962568e807f09ed808db911309d53c7b71bcd74d7c091763a3c01bb74693a91bff604d7d21d4d781459be249cbd4102c87e8204169b0b6bcc0fad684d7a375a0df0d30e9abfc070d6f641ec2a7382cd3483a1b4ca6b4081272bc1895aa4b30951299af76c0dbe00a9196e4ddd59316810c3fc1e9d236d17e28134186fdf1e244983a7bf7e53481e"
	result := encryptionUtils.Encrypt(Payload, SecretKey)

	if len(Payload) > len(result) {
		t.Errorf("Test Failed: %s, %s inputted, %s expected; %s received", Payload, SecretKey, expected, result)
	}
}

func TestDecrypt(t *testing.T) {
	encryptionUtils := utils.NewEncryptionUtils()
	encrypted := "5208ae703cc2fa0851347f55d3b76d3fd6035ee081d71a401e8bc92ebdc25d42440f62310bda60628537744ac03f200d78da9e61f1019ce02087b7ce6c976e7b2d8ad6aa978c532cea8f3e744cc6a5cafedc4ae6cd1b08a4ef75d6e37aa3c0c76954d16d57750be2980c2c91ac7ef0bbd0722abd59bf6be22493ea9b9759c3ff4d17f17ab670b0b6fc320e6de982313f1c4e74c0897f9f5a32d58e3e53050ae8fdbebba9009d0d1250fe34dcde1ebb42acbc22834a02f53889076140f0eb8db1"
	result := encryptionUtils.Decrypt(encrypted, SecretKey)

	if result.Cid != Cid {
		t.Errorf("Test Failed: %s, %s inputted, %s received", encrypted, SecretKey, result)
	}

	if result.Fp != Fp {
		t.Errorf("Test Failed: %s, %s inputted, %s received", encrypted, SecretKey, result)
	}
}
