package encrypt

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

//var aesCommonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
//var aesCommonKey128 = []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}

// 可以用的key做iv， 也可以自定义传入

// AesEncrypt Aes加密
func AesEncrypt(origData, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	aesIV := key[:blockSize]
	if len(iv) > 0 {
		aesIV = iv[0]
	}
	blockMode := cipher.NewCBCEncrypter(block, aesIV)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// AesDecrypt Aes解密
func AesDecrypt(crypted, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	aesIV := key[:blockSize]
	if len(iv) > 0 {
		aesIV = iv[0]
	}
	blockMode := cipher.NewCBCDecrypter(block, aesIV)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	if origData = PKCS5UnPadding(origData); len(origData) == 0 {
		return nil, errors.New("aes decrypt data failed,wrong aesKey or data")
	}

	return origData, nil
}

// PKCS5Padding 。。。
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding ...
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding { // Maybe slice bounds out of range,check origData and aeskey
		return []byte{}
	}
	return origData[:(length - unpadding)]
}

// EncodeBase64String ...
func EncodeBase64String(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// DecodeBase64String ...
func DecodeBase64String(src string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(src)
	return string(res), err
}

// ReadFile 读取文件
func ReadFile(filename string) ([]byte, error) {
	f, err := ioutil.ReadFile(filename) //A successful call returns err == nil
	return f, err
}

// RsaEncrypt Rsa加密 公钥加密
func RsaEncrypt(rawData []byte, publicKeyPath string) ([]byte, error) {

	publicKey, err := ioutil.ReadFile(publicKeyPath) //A successful call returns err == nil
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, rawData)
}

// RsaDecrypt Rsa解密 私钥解密
func RsaDecrypt(ciphertext []byte, privateKeyPath string) ([]byte, error) {
	privateKey, err := ioutil.ReadFile(privateKeyPath) //A successful call returns err == nil
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// SingWithRsa RSA签名...
func SignWithRsa(s string, signType crypto.Hash) (string, error) {
	h := crypto.Hash.New(signType)
	h.Write([]byte(s))
	hash := h.Sum(nil)

	// 调用签名 加密
	res, err := signature(hash, signType)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// signature 签名
func signature(hash []byte, signType crypto.Hash) (string, error) {
	privateKey, err := ioutil.ReadFile("private.pem") //A successful call returns err == nil
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("pem.Decode err")
	}

	private, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", errors.New("ParsePKCS8PrivateKey err")
	}

	// 签名算法
	signature, err := rsa.SignPKCS1v15(rand.Reader, private.(*rsa.PrivateKey),
		signType, hash)
	//signRet := fmt.Sprintf("%x", signature)
	// 16进制返回
	//return hex.EncodeToString(signature), err
	return string(signature), err

}

// ValidSignatureSha256 验签 对采用sha256签名的验签
func ValidSignatureSha256(originalData, signData string, pubKeyData []byte) error {
	block, _ := pem.Decode(pubKeyData)
	if block == nil {
		return errors.New("public key error")
	}
	//public, err := x509.ParsePKIXPublicKey ([]byte(pubKeyData))
	public, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	h := sha256.New()
	h.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(public.(*rsa.PublicKey), crypto.SHA256, h.Sum(nil), []byte(signData))
}
