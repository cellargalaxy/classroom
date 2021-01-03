package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/sirupsen/logrus"
)

const privateKeyType = "RSA PRIVATE KEY"
const publicKeyType = "RSA PUBLIC KEY"

var PublicKeyError = fmt.Errorf("public key error")
var PrivateKeyError = fmt.Errorf("private key error")

func CreateRsa() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("创建RSA私钥异常")
		return nil, nil, err
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  privateKeyType,
		Bytes: x509PrivateKey,
	}
	privateKeyBytes := pem.EncodeToMemory(block)

	publicKey := &privateKey.PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("创建RSA公钥异常")
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  publicKeyType,
		Bytes: x509PublicKey,
	}
	publicKeyBytes := pem.EncodeToMemory(block)
	return privateKeyBytes, publicKeyBytes, nil
}

// 签名验证
func VerifyRsa(publicKey []byte, data []byte, sign []byte) (bool, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		logrus.WithFields(logrus.Fields{}).Error("解析RSA公钥失败")
		return false, PublicKeyError
	}
	x509PublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("解析RSA公钥异常")
		return false, err
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(x509PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], sign)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("验证RSA签名异常")
		return false, err
	}
	return true, nil
}

// 公钥加密
func RsaEncrypt(publicKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		logrus.WithFields(logrus.Fields{}).Error("解析RSA公钥失败")
		return nil, PublicKeyError
	}
	x509PublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("解析RSA公钥异常")
		return nil, err
	}

	encrypt, err := rsa.EncryptPKCS1v15(rand.Reader, x509PublicKey.(*rsa.PublicKey), data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("RSA公钥加密异常")
		return nil, err
	}
	return encrypt, nil
}

// 私钥解密
func RsaDecrypt(privateKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		logrus.WithFields(logrus.Fields{}).Error("解析RSA私钥失败")
		return nil, PrivateKeyError
	}
	x509PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("解析RSA私钥异常")
		return nil, err
	}

	decrypt, err := rsa.DecryptPKCS1v15(rand.Reader, x509PrivateKey, data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("RSA私钥解密异常")
		return nil, err
	}
	return decrypt, nil
}
