package db

import (
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func InsertDataCreate(dataCreate *model.DataCreateAddModel) (*model.DataCreateAddModel, error) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&dataCreate.DataModel).Error
		if err != nil {
			return err
		}
		err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&dataCreate.DataCreateModel).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Warn("插入数据创建异常")
		return dataCreate, fmt.Errorf("插入创建数据异常: %+v", err)
	}
	logrus.WithFields(logrus.Fields{}).Info("插入数据创建完成")
	return dataCreate, nil
}
