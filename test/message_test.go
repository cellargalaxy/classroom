package test

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service"
	"github.com/cellargalaxy/classroom/util"
	"math/rand"
	"testing"
	"time"
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
	messageAdd.Data = randString(16)
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

func TestRandString(t *testing.T) {
	t.Logf("randString: %+v\n", randString(16))
	t.Logf("randString: %+v\n", randString(16))
	t.Logf("randString: %+v\n", randString(16))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}
