package service

import (
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service/db"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func AddMessage(message *model.MessageAdd) (*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("添加信息，请求为空")
		return nil, fmt.Errorf("添加信息，请求为空")
	}
	if message.PublishPublicKeyHash == "" {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("添加信息，发布公钥hash为空")
		return &message.Message, fmt.Errorf("添加信息，发布公钥hash为空")
	}

	var publishUser, parentUser *model.User
	var dataHashCount, createSignHashCount, publishSignHashCount int64
	var errGroup errgroup.Group
	errGroup.Go(func() error {
		var err error
		user := &model.User{}
		user.PublicKeyHash = message.PublishPublicKeyHash
		publishUser, err = GetUser(user)
		return err
	})
	errGroup.Go(func() error {
		var err error
		user := &model.User{}
		user.PublicKeyHash = message.ParentHash
		parentUser, err = GetUser(user)
		return err
	})
	errGroup.Go(func() error {
		var err error
		msg := &model.MessageInquiry{}
		msg.DataHash = message.ParentHash
		dataHashCount, err = ListMessageCount(msg)
		return err
	})
	errGroup.Go(func() error {
		var err error
		msg := &model.MessageInquiry{}
		msg.CreateSignHash = message.ParentHash
		createSignHashCount, err = ListMessageCount(msg)
		return err
	})
	errGroup.Go(func() error {
		var err error
		msg := &model.MessageInquiry{}
		msg.PublishSignHash = message.ParentHash
		publishSignHashCount, err = ListMessageCount(msg)
		return err
	})
	if err := errGroup.Wait(); err != nil {
		return &message.Message, err
	}
	if publishUser == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("添加信息，发布用户不存在")
		return &message.Message, fmt.Errorf("添加信息，发布用户不存在")
	}
	if dataHashCount+createSignHashCount+publishSignHashCount == 0 && parentUser == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("添加信息，父hash不存在")
		return &message.Message, fmt.Errorf("添加信息，父hash不存在")
	}

	return db.AddMessage(&message.Message, publishUser.PublicKey)
}

func ListMessageCount(message *model.MessageInquiry) (int64, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("查询信息数量，请求为空")
		return 0, fmt.Errorf("查询信息数量，请求为空")
	}
	count, err := db.ListMessageCount(message)
	return count, err
}

func ListMessage(message *model.MessageInquiry) ([]*model.Message, error) {
	if message == nil {
		logrus.WithFields(logrus.Fields{"message": util.ToJson(message)}).Error("查询信息，请求为空")
		return nil, fmt.Errorf("查询信息，请求为空")
	}
	list, err := db.ListMessage(message)
	return list, err
}
