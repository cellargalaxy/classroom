package util

import (
	"crypto/sha1"
	"encoding/base32"
	"github.com/sirupsen/logrus"
)

func CreateMagnetSha1(data []byte) ([]byte, error) {
	base32String := base32.StdEncoding.EncodeToString(data)
	hash := sha1.New()
	_, err := hash.Write([]byte(base32String))
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("创建磁力链接SHA1异常")
		return nil, err
	}
	bytes := hash.Sum(nil)
	sha1ed := string(bytes)
	logrus.WithFields(logrus.Fields{"sha1ed": sha1ed}).Info("创建磁力链接SHA1完成")
	return bytes, nil
}
