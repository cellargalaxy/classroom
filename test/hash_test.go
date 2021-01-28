package test

import (
	"github.com/cellargalaxy/classroom/util"
	"testing"
)

func TestSha256String(t *testing.T) {
	hashed := util.Sha256String("bjbbbkgbkxcj")
	t.Logf("hashed: %+v\n", hashed)
}
