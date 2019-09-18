package encrypt

import (
	"io/ioutil"
	"testing"
)
// TestRSAEncrypt ...
func BenchmarkRSAEncrypt(t *testing.B) {
	// encrtypt data
	//data := []byte("joshua")
	publicKeyPath := "../public.pem"

	for i:=0 ;i<100000;i++ {
		data, _ := ioutil.ReadFile("access.log")
		res, err := RsaEncrypt(data, publicKeyPath)
		if err != nil {
			t.Error(err)
			return
		}
		// decrypt data
		privatePath := "../private.pem"
		res, err = RsaDecrypt(res, privatePath)
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
	res, err := AesEncrypt(data, key)
	if err != nil {
		t.Error(err)
		return
	}
    //t.Log("encrypt data :",string(res))
	// decrypt data
	res, err = AesDecrypt(res, key)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Log("after decrypt:", string(res))
	t.Log("finish..")
}
