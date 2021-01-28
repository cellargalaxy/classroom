package service

import (
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service/db"
	"github.com/sirupsen/logrus"
)

func AddMessage(message *model.MessageAdd) (*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，请求为空")
		return nil, fmt.Errorf("添加信息，请求为空")
	}
	if message.PublishPrivateKeyHash == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，发布私钥hash为空")
		return &message.Message, fmt.Errorf("添加信息，发布私钥hash为空")
	}

	user := &model.User{}
	user.PrivateKeyHash = message.PublishPrivateKeyHash
	user, err := GetUser(user)
	if err != nil {
		return &message.Message, err
	}
	if user == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，发布用户不存在")
		return &message.Message, fmt.Errorf("添加信息，发布用户不存在")
	}

	if message.ParentHash == "" {
		message.ParentHash = user.PrivateKeyHash
	}

	return db.AddMessage(&message.Message, user.PublicKey)
}

func ListMessage(message *model.Message) ([]*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("查询信息，请求为空")
		return nil, fmt.Errorf("查询信息，请求为空")
	}
	list, err := db.ListMessage(message)
	return list, err
}
