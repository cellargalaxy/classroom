package service

import (
	"fmt"
	"github.com/cellargalaxy/classroom/conf"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service/db"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
)

func AddMessage(message *model.MessageAdd) (*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("添加信息，请求为空")
		return nil, fmt.Errorf("添加信息，请求为空")
	}
	if message.PublishPublicKeyHash == "" {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("添加信息，发布公钥hash为空")
		return &message.Message, fmt.Errorf("添加信息，发布公钥hash为空")
	}

	if message.ParentHash == "" {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("添加信息，父hash为空")
		return &message.Message, fmt.Errorf("添加信息，父hash为空")
	}
	bytes, err := util.HexDecode(message.ParentHash)
	if err != nil {
		return &message.Message, err
	}
	if len(bytes) > conf.GetMaxHashSize() {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("添加信息，父hash过大")
		return &message.Message, fmt.Errorf("添加信息，父hash过大")
	}

	publishUser := &model.User{}
	publishUser.PublicKeyHash = message.PublishPublicKeyHash
	publishUser, err = GetUser(publishUser)
	if err != nil {
		return &message.Message, err
	}

	return db.AddMessage(&message.Message, publishUser.PublicKey)
}

func ListMessageCount(message *model.MessageInquiry) (int64, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("查询信息数量，请求为空")
		return 0, fmt.Errorf("查询信息数量，请求为空")
	}
	count, err := db.ListMessageCount(message)
	return count, err
}

func ListMessage(message *model.MessageInquiry) ([]*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("查询信息，请求为空")
		return nil, fmt.Errorf("查询信息，请求为空")
	}
	list, err := db.ListMessage(message)
	return list, err
}
