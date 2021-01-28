package service

import (
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service/db"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
)

func AddUser(user *model.UserAdd) (*model.User, error) {
	if user == nil {
		logrus.WithFields(logrus.Fields{"user": util.ToJson(user)}).Error("添加用户，请求为空")
		return nil, fmt.Errorf("添加用户，请求为空")
	}
	return db.AddUser(&user.User, GetAdminPublicKeyHash(), user.Sign)
}

func GetUser(user *model.User) (*model.User, error) {
	if user == nil {
		logrus.WithFields(logrus.Fields{"user": util.ToJson(user)}).Error("查询用户，请求为空")
		return user, fmt.Errorf("查询用户，请求为空")
	}
	return db.GetUser(user)
}
