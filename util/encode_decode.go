package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
)

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("base64解码异常")
		return nil, fmt.Errorf("base64解码异常: %+v", err)
	}
	return bytes, nil
}

func HexEncode(data []byte) string {
	return fmt.Sprintf("%x", data)
}

func HexDecode(data string) ([]byte, error) {
	bytes, err := hex.DecodeString(data)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("Hex解码异常")
		return nil, fmt.Errorf("hex解码异常: %+v", err)
	}
	return bytes, nil
}
