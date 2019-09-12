package snlogic

import (
	"fmt"
	"testing"
)

func TestDecryptEncrypt(t *testing.T) {
	a := "{\"cid\":\"198a41ff-a10f-4cda-a2f3-a9ca80c0703b\",\"fp\":\"6d8cabd95987f8318b1fe01593d5c2a5.24700f9f1986800ab4fcc880530dd0ed\"}"
	key := "6EA4915349C0AAC6F6572DA4F6B00C42DAD33E75"
	b := "821cb59a6647f1edf597956243e564b00c120f8ac1674a153fbd707da0707fb236ea040d1665f3d294aa1943afbae1b26b2b795a127f883ec221c10c881a147bb8acb7e760cd6f04edc21c396ee1f6c9627d9bf1315c484a970ce8930c2ed1011af7e8569325c7edcdf70396f1abca8486eabec24567bf215d2e60382c40e5c42af075379dacdf959cb3fef74f9c9d15"
	py := "c94f22d1f3af43c34a3be7e0f522bf8282b112bdf52d540bedf3d27a2c1defa9c92e7267cd13c87b662635dd13541b398e3a712142a6647e4a1fbedc22755e6986ce90df974a41956bb02e790753a6ac825390876f988d4410c94758816e624e346a3e51d2b2945545b89c09d1292e05fac2c9fb9ae4dc69be8bfa551a1019ab"
	pyRes, _ := Decrypt(py, key)
	fmt.Print(pyRes)

	enc, _ := Encrypt(a, key)

	dec, _ := Decrypt(enc, key)
	dec1, _ := Decrypt(b, key)
	fmt.Print(dec)
	fmt.Print(dec1)
}
