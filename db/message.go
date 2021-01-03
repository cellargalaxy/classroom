package db

import (
	"errors"
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InsertMessage(message model.Message) (*model.Message, error) {
	err := db.Create(&message).Error
	return &message, err
}

func UpdateMessage(message model.Message) (*model.Message, error) {
	err := db.Updates(&message).Error
	return &message, err
}

func SelectMessage(message model.MessageInquiry) ([]*model.Message, error) {
	var where *gorm.DB
	if message.UserHash != "" {
		where = db.Where("user_hash = ?", message.UserHash)
	}
	if message.ParentHash != "" {
		where = db.Where("parent_hash = ?", message.ParentHash)
	}
	if message.Hash != "" {
		where = db.Where("hash = ?", message.Hash)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Warn("查询消息条件为空")
		return nil, fmt.Errorf("查询消息条件为空")
	}

	if message.StartCreatedAt.Unix() > 0 {
		where = where.Where("created_at >= ?", message.StartCreatedAt)
	}
	if message.EndCreatedAt.Unix() > 0 {
		where = where.Where("created_at < ?", message.EndCreatedAt)
	}
	if message.Offset >= 0 {
		where = where.Offset(message.Offset)
	}
	if message.Limit > 0 {
		where = where.Limit(message.Limit)
	}
	where = where.Where("deleted_at == null")

	var list []*model.Message
	err := where.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.WithFields(logrus.Fields{"message": message}).Warn("查询消息不存在")
		return nil, nil
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{"message": message, "err": err}).Error("查询消息异常")
		return nil, err
	}
	return list, err
}
