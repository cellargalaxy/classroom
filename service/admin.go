package service

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
)

const adminUserID = 1

var verifyData string

func init() {
	_, err := initAdminUser()
	if err != nil {
		logrus.Error("初始化管理员用户异常")
		panic(err)
	}
}

func initAdminUser() (*model.User, error) {
	user, err := getAdminUser()
	if err != nil {
		return user, err
	}
	if user != nil {
		logrus.Info("管理员用户已创建")
		return user, nil
	}

	privateKey, publicKey, err := util.CreateRsa()
	if err != nil {
		return nil, err
	}
	publicKeyHash := util.Sha256String(publicKey)
	setVerifyData(publicKeyHash)
	sign, err := util.RsaSignString(privateKey, GetVerifyData())
	if err != nil {
		return nil, err
	}

	var userAdd model.UserAdd
	userAdd.Id = adminUserID
	userAdd.PrivateKey = privateKey
	userAdd.PublicKey = publicKey
	userAdd.Sign = sign

	user, err = AddUser(&userAdd)
	return user, err
}

func setVerifyData(data string) {
	verifyData = data
}

func GetVerifyData() string {
	if verifyData != "" {
		return verifyData
	}
	user, err := getAdminUser()
	if err != nil {
		logrus.Error("查询管理员用户失败")
		panic(err)
	}
	if user.PublicKeyHash == "" {
		logrus.Error("查询管理员用户公钥HASH为空")
		panic(err)
	}
	setVerifyData(user.PublicKeyHash)
	return verifyData
}

func getAdminUser() (*model.User, error) {
	var user model.User
	user.Id = adminUserID
	return GetUser(&user)
}
