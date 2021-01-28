package test

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service"
	"github.com/cellargalaxy/classroom/util"
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
	messageAdd.Data = "abcde12345"
	messageAdd.DataHash = util.Sha256String(messageAdd.Data)
	createSign, err := util.RsaHashSignString(user.PrivateKey, messageAdd.DataHash)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	messageAdd.CreateSign = createSign
	messageAdd.CreatePublicKeyHash = user.PublicKeyHash

	publishSign, err := util.RsaHashSignString(user.PrivateKey, messageAdd.DataHash)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	messageAdd.PublishSign = publishSign
	messageAdd.PublishPrivateKeyHash = user.PrivateKeyHash

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

	message := &model.Message{}
	message.ParentHash = user.PrivateKeyHash
	t.Logf("message: %+v\n", message)
	list, err := service.ListMessage(message)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("len(list): %+v\n", len(list))
	for _, m := range list {
		t.Logf("message: %+v\n", util.ToJsonIndent(m))
	}
}
