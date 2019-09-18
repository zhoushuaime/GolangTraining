package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type MD5Client struct {
}

var MD5 = MD5Client{}

func (*MD5Client) Encrypt(str string) string {
	passwordBytes := []byte(str)
	sum := md5.Sum(passwordBytes)
	result := hex.EncodeToString(sum[:])
	return result
}

/*
salt
*/
func (*MD5Client) EncryptWithSalt(password string, salt string) string {
	passwordBytes := []byte(password)
	saltBytes := []byte(salt)
	hash := md5.New()
	hash.Write(passwordBytes)
	hash.Write(saltBytes)
	result := hash.Sum(nil)
	return hex.EncodeToString(result)
}
func main() {
	result1 := MD5.Encrypt("admin")
	result2 := MD5.EncryptWithSalt("test",
		strconv.FormatInt(time.Now().Unix(), 10)+"username")
	fmt.Println("result1:", result1)
	fmt.Println("result2:", result2)
	fmt.Println("==============================")
	r,err := strconv.ParseUint(result1,0,64)
	if err != nil {
		fmt.Println(err)
		return
	}
	rs := int64(r)
	if rs >=0 {
		fmt.Println("rs >= 0",rs)
	}

	if -rs >=0 {
		fmt.Println("-rs >= 0",rs)
	}
}
