package db

import (
	"github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gorm_logrus.New(),
	}

	var err error
	db, err = initSqlite(config)
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Error("初始化DB异常")
		panic(err)
	}
}

func initSqlite(config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("classroom.sqlite"), config)
}
