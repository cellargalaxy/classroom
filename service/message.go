package service

import (
	"fmt"
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service/db"
	"github.com/cellargalaxy/classroom/util"
	"github.com/sirupsen/logrus"
	"sync"
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
	var publishUserErr, parentUserErr error
	var dataHashCount, createSignHashCount, publishSignHashCount int64
	var dataHashErr, createSignHashErr, publishSignHashErr error
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	go func() {
		defer waitGroup.Done()
		user := &model.User{}
		user.PublicKeyHash = message.PublishPublicKeyHash
		publishUser, publishUserErr = GetUser(user)
	}()
	go func() {
		defer waitGroup.Done()
		user := &model.User{}
		user.PublicKeyHash = message.ParentHash
		parentUser, parentUserErr = GetUser(user)
	}()
	go func() {
		defer waitGroup.Done()
		msg := &model.MessageInquiry{}
		msg.DataHash = message.ParentHash
		dataHashCount, dataHashErr = ListMessageCount(msg)
	}()
	go func() {
		defer waitGroup.Done()
		msg := &model.MessageInquiry{}
		msg.CreateSignHash = message.ParentHash
		createSignHashCount, createSignHashErr = ListMessageCount(msg)
	}()
	go func() {
		defer waitGroup.Done()
		msg := &model.MessageInquiry{}
		msg.PublishSignHash = message.ParentHash
		publishSignHashCount, publishSignHashErr = ListMessageCount(msg)
	}()
	waitGroup.Wait()
	if publishUserErr != nil {
		return &message.Message, publishUserErr
	}
	if parentUserErr != nil {
		return &message.Message, parentUserErr
	}
	if dataHashErr != nil {
		return &message.Message, dataHashErr
	}
	if createSignHashErr != nil {
		return &message.Message, createSignHashErr
	}
	if publishSignHashErr != nil {
		return &message.Message, publishSignHashErr
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
