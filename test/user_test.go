package test

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service"
	"testing"
)

func TestGetUser(t *testing.T) {
	user := &model.User{}
	user.ID = 1
	user, err := service.GetUser(user)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("user: %+v\n", user)
}
