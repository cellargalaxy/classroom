package db

import (
	"errors"
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
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

func SelectMessageCount(message *model.MessageInquiry) (int64, error) {
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
	if message.CreateSignHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("create_sign_hash = ?", message.CreateSignHash)
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
	if message.PublishAtStart.Unix() > 0 {
		if where == nil {
			where = db
		}
		where = where.Where("publish_at >= ?", message.PublishAtStart)
	}
	if message.PublishAtEnd.Unix() > 0 {
		if where == nil {
			where = db
		}
		where = where.Where("publish_at < ?", message.PublishAtEnd)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Warn("查询消息条件为空")
		return 0, fmt.Errorf("查询消息条件为空")
	}

	var count int64
	err := where.Table(message.TableName()).Count(&count).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message), "err": err}).Error("查询消息数量异常")
		return 0, err
	}
	return count, err
}

func SelectMessage(message *model.MessageInquiry) ([]*model.Message, error) {
	var where *gorm.DB
	if message.Id > 0 {
		where = db.Where("id = ?", message.Id)
	}
	if message.ParentHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("parent_hash = ?", message.ParentHash)
	}
	if message.PublishAt.Unix() > 0 {
		if where == nil {
			where = db
		}
		where = where.Where("publish_at = ?", message.PublishAt)
	}
	if message.PublishAtStart.Unix() > 0 {
		if where == nil {
			where = db
		}
		where = where.Where("publish_at >= ?", message.PublishAtStart)
	}
	if message.PublishAtEnd.Unix() > 0 {
		if where == nil {
			where = db
		}
		where = where.Where("publish_at < ?", message.PublishAtEnd)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Warn("查询消息条件为空")
		return nil, fmt.Errorf("查询消息条件为空")
	}
	if message.Offset > 0 {
		where = where.Offset(message.Offset)
	}
	if message.Length > 0 {
		where = where.Limit(message.Length)
	}

	var list []*model.Message
	err := where.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Warn("查询消息不存在")
		return nil, nil
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message), "err": err}).Error("查询消息异常")
		return nil, err
	}
	return list, err
}
