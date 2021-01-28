package db

import (
	"errors"
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InsertUser(user *model.User) (*model.User, error) {
	err := db.Create(user).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("插入用户异常")
		return user, fmt.Errorf("插入用户异常: %+v", err)
	}
	return user, nil
}

func UpdateUser(user *model.User) (*model.User, error) {
	err := db.Updates(user).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("更新用户异常")
		return user, fmt.Errorf("更新用户异常: %+v", err)
	}
	return user, nil
}

func SelectUser(user *model.User) (*model.User, error) {
	var where *gorm.DB
	if user.Id > 0 {
		where = db.Where("id = ?", user.Id)
	}
	if user.PublicKeyHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("public_key_hash = ?", user.PublicKeyHash)
	}
	if user.PrivateKeyHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("private_key_hash = ?", user.PrivateKeyHash)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"user": user}).Warn("查询用户条件为空")
		return nil, fmt.Errorf("查询用户条件为空")
	}

	var selectUser model.User
	err := where.Take(&selectUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.WithFields(logrus.Fields{"user": user}).Warn("查询用户不存在")
		return nil, nil
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{"user": user, "err": err}).Error("查询用户异常")
		return nil, fmt.Errorf("查询用户异常: %+v", err)
	}
	return &selectUser, err
}
