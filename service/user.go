package service

import (
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service/db"
	"github.com/sirupsen/logrus"
)

func AddUser(user *model.UserAdd) (*model.User, error) {
	if user == nil {
		logrus.WithFields(logrus.Fields{"user": user}).Error("添加用户为空")
		return nil, fmt.Errorf("添加用户为空")
	}
	return db.AddUser(&user.User, GetVerifyData(), user.Sign)
}

func GetUser(user *model.User) (*model.User, error) {
	if user == nil {
		logrus.WithFields(logrus.Fields{"user": user}).Error("查询用户为空")
		return user, fmt.Errorf("查询用户为空")
	}
	return db.GetUser(user)
}
