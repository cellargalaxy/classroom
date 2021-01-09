package test

import (
	"github.com/cellargalaxy/classroom/util"
	"testing"
)

func TestSha256String(t *testing.T) {
	hashed, err := util.Sha256String("bjbbbkgbkxcj")
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("hashed: %+v\n", hashed)
}
