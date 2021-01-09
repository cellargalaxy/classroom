package util

import (
	"crypto/sha256"
	"fmt"
	"github.com/sirupsen/logrus"
)

func Sha256String(data string) (string, error) {
	hashed, err := Sha256([]byte(data))
	if err != nil {
		return "", err
	}
	return Hex(hashed), nil
}

func Sha256(data []byte) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("sha256异常")
		return nil, err
	}
	return hash.Sum(nil), nil
}

func Hex(data []byte) string {
	return fmt.Sprintf("%x", data)
}
