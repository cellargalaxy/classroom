package service

import (
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service/db"
	"github.com/sirupsen/logrus"
)

func AddMessage(message *model.MessageAdd) (*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息请求为空")
		return nil, fmt.Errorf("添加信息请求为空")
	}
	if message.UserHash == "" {
		logrus.WithFields(logrus.Fields{"message": message}).Error("添加信息，用户Hash为空")
		return &message.Message, fmt.Errorf("添加信息，用户Hash为空")
	}

	user := &model.User{}
	user.PublicKeyHash = message.UserHash
	user, err := GetUser(user)
	if err != nil {
		return &message.Message, err
	}

	if message.ParentHash == "" {
		message.ParentHash = user.PublicKeyHash
	}

	return db.AddMessage(&message.Message, message.Sign, user.PublicKey)
}

func ListMessage(message *model.MessageInquiry) ([]*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Error("查询信息请求为空")
		return nil, fmt.Errorf("查询信息请求为空")
	}
	list, err := db.ListMessage(message)
	return list, err
}
