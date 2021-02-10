package db

import (
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/sirupsen/logrus"
)

func InsertDataPublish(dataPublish *model.DataPublishModel) (*model.DataPublishModel, error) {
	err := db.Create(dataPublish).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("插入数据发布异常")
		return dataPublish, fmt.Errorf("插入数据发布异常: %+v", err)
	}
	logrus.WithFields(logrus.Fields{}).Info("插入数据发布完成")
	return dataPublish, nil
}
