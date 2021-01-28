package test

import (
	"github.com/cellargalaxy/classroom/model"
	"github.com/cellargalaxy/classroom/service"
	"github.com/cellargalaxy/classroom/util"
	"testing"
)

func TestGetUser(t *testing.T) {
	user := &model.User{}
	user.Id = 1
	user, err := service.GetUser(user)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("user: %+v\n", util.ToJsonIndent(user))
}
