package db

import (
	"crypto/sha1"
	"fmt"
	"github.com/cellargalaxy/classroom/db"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
)

func AddUser(userAdd model.UserAdd, signMessage []byte) (*model.User, error) {
	if userAdd.PublicKey == "" {
		logrus.WithFields(logrus.Fields{"userAdd": userAdd}).Error("添加用户，公钥为空")
		return &userAdd.User, fmt.Errorf("添加用户，公钥为空")
	}
	if userAdd.Sign == "" {
		logrus.WithFields(logrus.Fields{"userAdd": userAdd}).Error("添加用户，签名为空")
		return &userAdd.User, fmt.Errorf("添加用户，签名为空")
	}

	result, err := util.VerifyRsa(signMessage, []byte(userAdd.Sign), []byte(userAdd.PublicKey))
	if err != nil {
		return nil, err
	}
	if !result {
		logrus.WithFields(logrus.Fields{"userAdd": userAdd}).Error("添加用户，签名认证失败")
		return &userAdd.User, fmt.Errorf("添加用户，签名认证失败")
	}

	hash := sha1.New()
	hash.Write([]byte(userAdd.PublicKey))
	bytes := hash.Sum(nil)
	userAdd.PublicKeyHash = string(bytes)

	user, err := db.InsertUser(userAdd.User)
	return user, err
}

func GetUser(user model.User) (*model.User, error) {
	return db.SelectUser(user)
}
