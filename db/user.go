package db

import (
	"errors"
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InsertUser(user model.User) (*model.User, error) {
	err := db.Create(&user).Error
	return &user, err
}

func UpdateUser(user model.User) (*model.User, error) {
	err := db.Updates(&user).Error
	return &user, err
}

func SelectUser(user model.User) (*model.User, error) {
	var where *gorm.DB
	if user.ID > 0 {
		where = db.Where("id = ?", user.ID)
	}
	if user.PublicKeyHash != "" {
		where = db.Where("public_key_hash = ?", user.PublicKeyHash)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"user": user}).Warn("查询用户条件为空")
		return nil, fmt.Errorf("查询用户条件为空")
	}

	where = where.Where("deleted_at == null")

	err := where.Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.WithFields(logrus.Fields{"user": user}).Warn("查询用户不存在")
		return nil, nil
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{"user": user, "err": err}).Error("查询用户异常")
		return nil, err
	}
	return &user, err
}
