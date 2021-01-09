package test

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service"
	"github.com/cellargalaxy/classroom/util"
	"testing"
)

func TestAddMessage(t *testing.T) {
	user := &model.User{}
	user.ID = 1
	user, err := service.GetUser(user)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}

	messageAdd := &model.MessageAdd{}
	messageAdd.UserHash = user.PublicKeyHash
	messageAdd.Data = util.Base64Encode([]byte("abcde12345"))
	messageAdd.DataType = "text/plain"
	hashed, err := util.Sha256String(messageAdd.Data)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	messageAdd.Hash = hashed
	sign, err := util.RsaSignString(user.PrivateKey, messageAdd.Hash)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	messageAdd.Sign = sign

	message, err := service.AddMessage(messageAdd)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("message: %+v\n", message)
}

func TestListMessage(t *testing.T) {
	user := &model.User{}
	user.ID = 1
	user, err := service.GetUser(user)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}

	message := &model.MessageInquiry{}
	message.UserHash = user.PublicKeyHash
	t.Logf("message: %+v\n", message)
	list, err := service.ListMessage(message)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("len(list): %+v\n", len(list))
	for _, m := range list {
		t.Logf("message: %+v\n", m)
	}
}
