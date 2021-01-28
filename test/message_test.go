package test

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service"
	"github.com/cellargalaxy/classroom/util"
	"math/rand"
	"testing"
)

func TestAddMessage(t *testing.T) {
	user := &model.User{}
	user.Id = 1
	user, err := service.GetUser(user)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}

	messageAdd := &model.MessageAdd{}
	messageAdd.DataType = "text/plain"
	messageAdd.Data = randStringRunes(16)
	messageAdd.DataHash = util.Sha256String(messageAdd.Data)

	createSign, err := util.RsaHashSignString(user.PrivateKey, messageAdd.DataHash)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	messageAdd.CreateSign = createSign
	messageAdd.CreatePublicKeyHash = user.PublicKeyHash
	messageAdd.CreateSignHash = util.Sha256String(messageAdd.CreateSign)

	publishSign, err := util.RsaHashSignString(user.PrivateKey, messageAdd.CreateSignHash)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	messageAdd.PublishSign = publishSign
	messageAdd.PublishPublicKeyHash = user.PublicKeyHash
	messageAdd.PublishSignHash = util.Sha256String(messageAdd.PublishSign)

	messageAdd.ParentHash = user.PublicKeyHash

	message, err := service.AddMessage(messageAdd)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("message: %+v\n", util.ToJsonIndent(message))
}

func TestListMessage(t *testing.T) {
	user := &model.User{}
	user.Id = 1
	user, err := service.GetUser(user)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}

	message := &model.MessageInquiry{}
	message.ParentHash = user.PublicKeyHash
	t.Logf("message: %+v\n", util.ToJsonIndent(message))
	list, err := service.ListMessage(message)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("len(list): %+v\n", len(list))
	for _, m := range list {
		t.Logf("message.Data: %+v\n", m.Data)
		t.Logf("message.CreateSign: %+v\n", m.CreateSign)
		t.Logf("message.PublishSign: %+v\n", m.PublishSign)
		t.Logf("message: %+v\n", util.ToJsonIndent(m))
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
