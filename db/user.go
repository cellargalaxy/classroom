package db

import (
	"errors"
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func InsertUser(user *model.UserModel) (*model.UserModel, error) {
	err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(user).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("插入用户异常")
		return user, fmt.Errorf("插入用户异常: %+v", err)
	}
	logrus.WithFields(logrus.Fields{}).Info("插入用户完成")
	return user, nil
}

func UpdateUser(user *model.UserModel) (*model.UserModel, error) {
	err := db.Updates(user).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("更新用户异常")
		return user, fmt.Errorf("更新用户异常: %+v", err)
	}
	logrus.WithFields(logrus.Fields{}).Info("更新用户完成")
	return user, nil
}

func SelectUser(user *model.UserModel) (*model.UserModel, error) {
	var where *gorm.DB
	if user.ID > 0 {
		where = db.Where("id = ?", user.ID)
	}
	if user.PublicKeyHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("public_key_hash = ?", user.PublicKeyHash)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"user": util.ToJson(user)}).Warn("查询用户条件为空")
		return nil, fmt.Errorf("查询用户条件为空")
	}

	var object model.UserModel
	err := where.Take(&object).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.WithFields(logrus.Fields{"user": util.ToJson(user)}).Warn("查询用户不存在")
		return nil, nil
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{"user": util.ToJson(user), "err": err}).Error("查询用户异常")
		return nil, fmt.Errorf("查询用户异常: %+v", err)
	}
	logrus.WithFields(logrus.Fields{"user": util.ToJson(object)}).Info("查询用户完成")
	return &object, err
}
