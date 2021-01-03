package service

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service/db"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
)

const adminUserID = 1

var signMessage []byte

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
	publicKeyHash, err := util.CreateMagnetSha1(publicKey)
	if err != nil {
		return nil, err
	}
	setSignMessage(publicKeyHash)
	sign, err := util.RsaEncrypt(publicKey, GetSignMessage())
	if err != nil {
		return nil, err
	}

	var userAdd model.UserAdd
	userAdd.ID = adminUserID
	userAdd.PrivateKey = string(privateKey)
	userAdd.PublicKey = string(publicKey)
	userAdd.PublicKeyHash = string(publicKeyHash)
	userAdd.Sign = string(sign)

	user, err = db.AddUser(userAdd, GetSignMessage())
	return user, err
}

func setSignMessage(sign []byte) {
	signMessage = sign
}

func GetSignMessage() []byte {
	if signMessage != nil {
		return signMessage
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
	setSignMessage([]byte(user.PublicKeyHash))
	return signMessage
}

func getAdminUser() (*model.User, error) {
	var user model.User
	user.ID = adminUserID
	userP, err := db.GetUser(user)
	return userP, err
}
