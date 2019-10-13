package encrypt_test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"zhoushuai.com/GolangTraining/encrypt"
)

var aesCommonKey128 = []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
var aesCommonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// TestAesEncrypt ...
func TestAesEncrypt(t *testing.T) {
	originalText := "admin"
	fmt.Println("originalText:", originalText)
	mytext := []byte(originalText)
	cryptoText, err := encrypt.AesEncrypt(mytext, aesCommonKey128, aesCommonIV)
	if !assert.Nil(t, err, fmt.Sprintf("aes encrypt error:%v", err)) {
		return
	}
	fmt.Println("cryptoText hex:", hex.EncodeToString(cryptoText))
	fmt.Println("cryptoText base64:", base64.StdEncoding.EncodeToString(cryptoText))

}

// TestAesDecryption ...
func TestAesDecryption(t *testing.T) {
	originalText := "QH7yT5TDkOGNRmSZC6htpQ=="
	//originalText := "dCjlRtxsBp66W0Hbe2TD8g=="
	byteData, err := base64.StdEncoding.DecodeString(originalText)
	if !assert.Nil(t, err, fmt.Sprintf("encrypted data base64 decode error:%v", err)) {
		return
	}
	decryptedText, err := encrypt.AesDecrypt(byteData, aesCommonKey128, aesCommonIV)
	if !assert.Nil(t, err, fmt.Sprintf("aes decrypt error:%v", err)) {
		return
	}
	expected := "admin"
	acutal := string(decryptedText)
	fmt.Println("decryptedText:", string(decryptedText))
	assert.Equal(t, expected, acutal, "not match with actual data")
}

// TestRSAEncrypt ...
func BenchmarkRSAEncrypt(t *testing.B) {
	// encrtypt data
	//data := []byte("joshua")
	publicKeyPath := "../public.pem"

	for i := 0; i < 100000; i++ {
		data, _ := ioutil.ReadFile("access.log")
		res, err := encrypt.RsaEncrypt(data, publicKeyPath)
		if err != nil {
			t.Error(err)
			return
		}
		// decrypt data
		privatePath := "../private.pem"
		res, err = encrypt.RsaDecrypt(res, privatePath)
		if err != nil {
			t.Error(err)
			return
		}
	}

	//t.Log("after decrypt:", string(res))
	t.Log("finish")
}

// BenchmarkAesEncrypt ...
func BenchmarkAesEncrypt(t *testing.B) {
	//data := []byte("zhoushuai")
	key := []byte("aeTh6ae2eu2phoob")
	// encrtypt data
	data, _ := ioutil.ReadFile("stderr.log")
	res, err := encrypt.AesEncrypt(data, key)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Log("encrypt data :",string(res))
	// decrypt data
	res, err = encrypt.AesDecrypt(res, key)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Log("after decrypt:", string(res))
	t.Log("finish..")
}
