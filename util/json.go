package util

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func ToJson(x interface{}) string {
	bytes, err := json.Marshal(x)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("系列化json异常")
	}
	return string(bytes)
}

func ToJsonIndent(x interface{}) string {
	bytes, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("系列化json异常")
	}
	return string(bytes)
}
