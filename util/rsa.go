package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/sirupsen/logrus"
)

const privateKeyType = "RSA PRIVATE KEY"
const publicKeyType = "RSA PUBLIC KEY"

func CreateRsa() (string, string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("创建RSA私钥异常")
		return "", "", fmt.Errorf("创建RSA私钥异常: %+v", err)
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
		return "", "", fmt.Errorf("创建RSA公钥异常: %+v", err)
	}
	block = &pem.Block{
		Type:  publicKeyType,
		Bytes: x509PublicKey,
	}
	publicKeyBytes := pem.EncodeToMemory(block)
	return string(privateKeyBytes), string(publicKeyBytes), nil
}

// 公钥加密
func RsaEncryptString(publicKey, data string) (string, error) {
	encrypt, err := RsaEncrypt(publicKey, []byte(data))
	if err != nil {
		return "", err
	}
	return Base64Encode(encrypt), nil
}

func RsaEncrypt(publicKey string, data []byte) ([]byte, error) {
	key, err := ParsePublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	encrypt, err := rsa.EncryptPKCS1v15(rand.Reader, key, data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("RSA公钥加密异常")
		return nil, fmt.Errorf("RSA公钥加密异常: %+v", err)
	}
	return encrypt, nil
}

// 私钥解密
func RsaDecryptString(privateKey, data string) (string, error) {
	bytes, err := Base64Decode(data)
	if err != nil {
		return "", err
	}
	decrypt, err := RsaDecrypt(privateKey, bytes)
	if err != nil {
		return "", err
	}
	return string(decrypt), nil
}

func RsaDecrypt(privateKey string, data []byte) ([]byte, error) {
	key, err := ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	decrypt, err := rsa.DecryptPKCS1v15(rand.Reader, key, data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("RSA私钥解密异常")
		return nil, fmt.Errorf("RSA私钥解密异常: %+v", err)
	}
	return decrypt, nil
}

// 私钥签名
func RsaSignString(privateKey, data string) (string, error) {
	sign, err := RsaSign(privateKey, []byte(data))
	if err != nil {
		return "", err
	}
	return Base64Encode(sign), nil
}

func RsaSign(privateKey string, data []byte) ([]byte, error) {
	hashed := Sha256(data)
	return RsaHashSign(privateKey, hashed)
}

func RsaHashSignString(privateKey string, hashed string) (string, error) {
	hashBytes, err := HexDecode(hashed)
	if err != nil {
		return "", err
	}
	signBytes, err := RsaHashSign(privateKey, hashBytes)
	if err != nil {
		return "", err
	}
	return Base64Encode(signBytes), nil
}

func RsaHashSign(privateKey string, hashed []byte) ([]byte, error) {
	key, err := ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	sign, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("私钥签名异常")
		return nil, fmt.Errorf("私钥签名异常: %+v", err)
	}
	return sign, nil
}

func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		logrus.WithFields(logrus.Fields{}).Error("解析RSA私钥失败")
		return nil, fmt.Errorf("解析RSA私钥失败")
	}
	x509PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("解析RSA私钥异常")
		return nil, fmt.Errorf("解析RSA私钥异常: %+v", err)
	}
	return x509PrivateKey, nil
}

// 公钥验证
func RsaVerifyString(publicKey, data, sign string) (bool, error) {
	signBytes, err := Base64Decode(sign)
	if err != nil {
		return false, err
	}
	return RsaVerify(publicKey, []byte(data), signBytes)
}

func RsaVerify(publicKey string, data, sign []byte) (bool, error) {
	hashBytes := Sha256(data)
	return RsaHashVerify(publicKey, hashBytes, sign)
}

func RsaHashVerifyString(publicKey string, hashed, sign string) (bool, error) {
	hashBytes, err := HexDecode(hashed)
	if err != nil {
		return false, err
	}
	signBytes, err := Base64Decode(sign)
	if err != nil {
		return false, err
	}
	return RsaHashVerify(publicKey, hashBytes, signBytes)
}

func RsaHashVerify(publicKey string, hashed, sign []byte) (bool, error) {
	key, err := ParsePublicKey(publicKey)
	if err != nil {
		return false, err
	}

	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed, sign)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("验证RSA签名异常")
		return false, fmt.Errorf("验证RSA签名异常: %+v", err)
	}
	return true, nil
}

func ParsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		logrus.WithFields(logrus.Fields{}).Error("解析RSA公钥失败")
		return nil, fmt.Errorf("解析RSA公钥失败")
	}
	x509PublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("解析RSA公钥异常")
		return nil, fmt.Errorf("解析RSA公钥异常: %+v", err)
	}
	key, ok := x509PublicKey.(*rsa.PublicKey)
	if !ok {
		logrus.WithFields(logrus.Fields{}).Error("解析RSA公钥转型失败")
		return nil, fmt.Errorf("解析RSA公钥转型失败: %+v", err)
	}
	return key, nil
}

func CheckPublicPrivateKey(publicKey, privateKey string) (bool, error) {
	sign, err := RsaSignString(privateKey, "")
	if err != nil {
		return false, err
	}
	return RsaVerifyString(publicKey, "", sign)
}
