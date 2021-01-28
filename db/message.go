package db

import (
	"errors"
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InsertMessage(message *model.Message) (*model.Message, error) {
	err := db.Create(message).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("插入消息异常")
		return message, fmt.Errorf("插入消息异常: %+v", err)
	}
	return message, err
}

func UpdateMessage(message *model.Message) (*model.Message, error) {
	err := db.Updates(message).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("更新消息异常")
		return message, fmt.Errorf("更新消息异常: %+v", err)
	}
	return message, err
}

func SelectMessage(message *model.Message) ([]*model.Message, error) {
	var where *gorm.DB
	if message.Id > 0 {
		where = db.Where("id = ?", message.Id)
	}
	if message.DataHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("data_hash = ?", message.DataHash)
	}
	if message.PublishSignHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("publish_sign_hash = ?", message.PublishSignHash)
	}
	if message.PublishAt.Unix() > 0 {
		if where == nil {
			where = db
		}
		where = where.Where("publish_at = ?", message.PublishAt)
	}
	if message.ParentHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("parent_hash = ?", message.ParentHash)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"message": message}).Warn("查询消息条件为空")
		return nil, fmt.Errorf("查询消息条件为空")
	}

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
