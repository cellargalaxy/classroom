package db

import (
	"fmt"
	"github.com/cellargalaxy/classroom/db"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
)

func AddUser(user *model.User, verifyData, sign string) (*model.User, error) {
	if user == nil {
		logrus.WithFields(logrus.Fields{"user": user}).Error("添加用户请求为空")
		return user, fmt.Errorf("添加用户请求为空")
	}
	if user.PublicKey == "" {
		logrus.WithFields(logrus.Fields{"user": user}).Error("添加用户，公钥为空")
		return user, fmt.Errorf("添加用户，公钥为空")
	}
	if sign == "" {
		logrus.WithFields(logrus.Fields{"user": user}).Error("添加用户，签名为空")
		return user, fmt.Errorf("添加用户，签名为空")
	}

	result, err := util.RsaVerifyString(user.PublicKey, verifyData, sign)
	if err != nil {
		return user, err
	}
	if !result {
		logrus.WithFields(logrus.Fields{"user": user, "verifyData": verifyData, "sign": sign}).Error("添加用户，签名认证失败")
		return user, fmt.Errorf("添加用户，签名认证失败")
	}

	hashed, err := util.Sha256String(user.PublicKey)
	if err != nil {
		return user, err
	}
	user.PublicKeyHash = hashed

	user, err = db.InsertUser(user)
	return user, err
}

func GetUser(user *model.User) (*model.User, error) {
	if user == nil {
		logrus.WithFields(logrus.Fields{"user": user}).Error("查询用户请求为空")
		return user, fmt.Errorf("查询用户请求为空")
	}
	return db.SelectUser(user)
}
