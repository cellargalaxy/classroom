package db

import (
	"fmt"
	"github.com/cellargalaxy/classroom/conf"
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
	size := len([]byte(message.Data))
	if size > conf.GetMaxDataSize() {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，数据大小超过限制")
		return message, fmt.Errorf("添加信息，数据大小超过限制")
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
	createSignHash := util.Sha256String(message.CreateSign)
	if message.CreateSignHash == "" {
		message.CreateSignHash = createSignHash
	}
	if message.CreateSignHash != createSignHash {
		logrus.WithFields(logrus.Fields{"message": message, "createSignHash": createSignHash}).Error("添加信息，创建签名hash校验失败")
		return message, fmt.Errorf("添加信息，创建签名hash校验失败")
	}

	if message.PublishSign == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，发布签名为空")
		return message, fmt.Errorf("添加信息，发布签名为空")
	}
	if message.PublishPublicKeyHash == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，发布公钥hash为空")
		return message, fmt.Errorf("添加信息，发布公钥hash为空")
	}
	publishSignHash := util.Sha256String(message.PublishSign)
	if message.PublishSignHash == "" {
		message.PublishSignHash = publishSignHash
	}
	if message.PublishSignHash != publishSignHash {
		logrus.WithFields(logrus.Fields{"message": message, "publishSignHash": publishSignHash}).Error("添加信息，发布签名hash校验失败")
		return message, fmt.Errorf("添加信息，发布签名hash校验失败")
	}
	if message.ParentHash == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，发布hash为空")
		return message, fmt.Errorf("添加信息，发布hash为空")
	}
	message.PublishAt = time.Now()

	if publicKey == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，公钥为空")
		return message, fmt.Errorf("添加信息，公钥为空")
	}

	result, err := util.RsaHashVerifyString(publicKey, message.CreateSignHash, message.PublishSign)
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

func ListMessageCount(message *model.MessageInquiry) (int64, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("查询信息数量，请求为空")
		return 0, fmt.Errorf("查询信息数量，请求为空")
	}
	count, err := db.SelectMessageCount(message)
	return count, err
}

func ListMessage(message *model.MessageInquiry) ([]*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("查询信息，请求为空")
		return nil, fmt.Errorf("查询信息，请求为空")
	}
	list, err := db.SelectMessage(message)
	return list, err
}
