package service

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
)

const adminUserID = 1

var adminPublicKeyHash string

func init() {
	admin, err := getAdminUser()
	if err != nil {
		logrus.Error("初始化管理员用户异常")
		panic(err)
	}
	setAdminPublicKeyHash(admin.PublicKeyHash)
}

func setAdminPublicKeyHash(hashed string) {
	adminPublicKeyHash = hashed
}

func GetAdminPublicKeyHash() string {
	return adminPublicKeyHash
}

func getAdminUser() (*model.User, error) {
	user := &model.User{}
	user.Id = adminUserID
	user, err := GetUser(user)
	if err != nil {
		return user, err
	}
	if user != nil {
		return user, nil
	}

	privateKey, publicKey, err := util.CreateRsa()
	if err != nil {
		return nil, err
	}
	publicKeyHash := util.Sha256String(publicKey)
	setAdminPublicKeyHash(publicKeyHash)
	sign, err := util.RsaHashSignString(privateKey, GetAdminPublicKeyHash())
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
