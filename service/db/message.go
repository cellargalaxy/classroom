package db

import (
	"fmt"
	"github.com/cellargalaxy/classroom/db"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
	"time"
)

func AddMessage(message *model.Message, publicKey string) (*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，请求为空")
		return message, fmt.Errorf("添加信息，请求为空")
	}
	if message.DataType == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，数据类型为空")
		return message, fmt.Errorf("添加信息，数据类型为空")
	}
	dataHash := util.Sha256String(message.Data)
	if message.DataHash == "" {
		message.DataHash = dataHash
	}
	if message.DataHash != dataHash {
		logrus.WithFields(logrus.Fields{"message": message, "dataHash": dataHash}).Error("添加信息，数据hash校验失败")
		return message, fmt.Errorf("添加信息，数据hash校验失败")
	}
	if message.CreateSign == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，创建签名为空")
		return message, fmt.Errorf("添加信息，创建签名为空")
	}
	if message.CreatePublicKeyHash == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，创建公钥hash为空")
		return message, fmt.Errorf("添加信息，创建公钥hash为空")
	}
	if message.PublishSign == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，发布签名为空")
		return message, fmt.Errorf("添加信息，发布签名为空")
	}
	if message.PublishPrivateKeyHash == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，发布私钥hash为空")
		return message, fmt.Errorf("添加信息，发布私钥hash为空")
	}
	message.PublishSignHash = util.Sha256String(message.PublishSign)
	if message.ParentHash == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，父hash为空")
		return message, fmt.Errorf("添加信息，父hash为空")
	}
	message.PublishAt = time.Now()

	if publicKey == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，公钥为空")
		return message, fmt.Errorf("添加信息，公钥为空")
	}

	result, err := util.RsaHashVerifyString(publicKey, message.DataHash, message.PublishSign)
	if err != nil {
		return message, err
	}
	if !result {
		logrus.WithFields(logrus.Fields{"message": message, "dataHash": message.DataHash, "publicKey": publicKey}).Error("添加信息，签名认证失败")
		return message, fmt.Errorf("添加信息，签名认证失败")
	}

	message.Id = 0
	var zeroTime time.Time
	message.CreatedAt = zeroTime
	message.UpdatedAt = zeroTime

	message, err = db.InsertMessage(message)
	return message, err
}

func ListMessage(message *model.Message) ([]*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("查询信息，请求为空")
		return nil, fmt.Errorf("查询信息，请求为空")
	}
	list, err := db.SelectMessage(message)
	return list, err
}
