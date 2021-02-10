package db

import (
	"errors"
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SelectOneData(data *model.DataInquiry) (*model.DataModel, error) {
	var where *gorm.DB
	if data.ID > 0 {
		where = db.Where("id = ?", data.ID)
	}
	if data.DataHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("data_hash = ?", data.DataHash)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"data": util.ToJson(data)}).Warn("查询数据条件为空")
		return nil, fmt.Errorf("查询数据条件为空")
	}

	var object model.DataModel
	err := where.Take(&object).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.WithFields(logrus.Fields{"data": util.ToJson(data)}).Warn("查询数据不存在")
		return nil, nil
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{"data": util.ToJson(data), "err": err}).Error("查询数据异常")
		return nil, fmt.Errorf("查询数据异常: %+v", err)
	}
	logrus.WithFields(logrus.Fields{"data": util.ToJson(object)}).Info("查询数据完成")
	return &object, err
}

func SelectSomeData(data *model.DataInquiry) ([]*model.DataModel, error) {
	var where *gorm.DB
	if data.ID > 0 {
		where = db.Where("id = ?", data.ID)
	}
	if data.DataHash != "" {
		if where == nil {
			where = db
		}
		where = where.Where("data_hash = ?", data.DataHash)
	}
	if len(data.DataHashes) > 0 {
		if where == nil {
			where = db
		}
		where = where.Where("data_hash in (?)", data.DataHashes)
	}
	if where == nil {
		logrus.WithFields(logrus.Fields{"data": util.ToJson(data)}).Warn("查询数据条件为空")
		return nil, fmt.Errorf("查询数据条件为空")
	}

	var list []*model.DataModel
	err := where.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logrus.WithFields(logrus.Fields{"data": util.ToJson(data)}).Warn("查询数据不存在")
		return nil, nil
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{"data": util.ToJson(data), "err": err}).Error("查询数据异常")
		return nil, fmt.Errorf("查询数据异常: %+v", err)
	}
	logrus.WithFields(logrus.Fields{"data": len(list)}).Info("查询数据完成")
	return list, err
}
