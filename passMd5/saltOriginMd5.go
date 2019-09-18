package main


import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type MD5Client struct {
}

var MD5 = MD5Client{}

func (this *MD5Client) Encrypt(plantext []byte) []byte {
	result := md5.Sum(plantext)
	return result[:]
}

/*
给要加密的信息加盐
*/
func (this *MD5Client) EncryptWithSalt(plantext []byte, salt []byte) []byte {
	hash := md5.New()
	hash.Write(plantext)
	hash.Write(salt)
	return hash.Sum(nil)
}

func main() {
	sum := MD5.Encrypt([]byte(`admin`))
	sumStr := hex.EncodeToString(sum)
	fmt.Println(sumStr)
}
