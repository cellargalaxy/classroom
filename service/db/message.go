package db

import (
	"fmt"
	"github.com/cellargalaxy/classroom/db"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
)

func AddMessage(message *model.Message, sign, publicKey string) (*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息请求为空")
		return message, fmt.Errorf("添加信息请求为空")
	}
	if message.Hash == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，Hash为空")
		return message, fmt.Errorf("添加信息，Hash为空")
	}
	if sign == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，签名为空")
		return message, fmt.Errorf("添加信息，签名为空")
	}
	if publicKey == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，公钥为空")
		return message, fmt.Errorf("添加信息，公钥为空")
	}

	result, err := util.RsaVerifyString(publicKey, message.Hash, sign)
	if err != nil {
		return message, err
	}
	if !result {
		logrus.WithFields(logrus.Fields{"message": message, "sign": sign, "publicKey": publicKey}).Error("添加信息，签名认证失败")
		return message, fmt.Errorf("添加信息，签名认证失败")
	}

	hashed, err := util.Sha256String(message.Data)
	if err != nil {
		return message, err
	}
	if hashed != message.Hash {
		logrus.WithFields(logrus.Fields{"message": message, "hashed": hashed}).Error("添加信息，数据HASH校验失败")
		return message, fmt.Errorf("添加信息，数据HASH校验失败")
	}

	message, err = db.InsertMessage(message)
	return message, err
}

func ListMessage(message *model.MessageInquiry) ([]*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("查询信息请求为空")
		return nil, fmt.Errorf("查询信息请求为空")
	}
	list, err := db.SelectMessage(message)
	return list, err
}
