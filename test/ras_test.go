package test

import (
	"github.com/cellargalaxy/classroom/util"
	"testing"
)

func TestRasCrypt(t *testing.T) {
	privateKey, publicKey, err := util.CreateRsa()
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("privateKey: %+v\n", privateKey)
	t.Logf("publicKey: %+v\n", publicKey)

	data := "abcde12345"
	encrypt, err := util.RsaEncryptString(publicKey, data)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("encrypt: %+v\n", encrypt)
	decrypt, err := util.RsaDecryptString(privateKey, encrypt)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("decrypt: %+v\n", decrypt)
	if decrypt != data {
		t.Errorf("加密解密失败\n")
		return
	}
}

func TestRasSign(t *testing.T) {
	privateKey, publicKey, err := util.CreateRsa()
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("privateKey: %+v\n", privateKey)
	t.Logf("publicKey: %+v\n", publicKey)

	data := "abcde12345"
	sign, err := util.RsaSignString(privateKey, data)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("sign: %+v\n", sign)
	verify, err := util.RsaVerifyString(publicKey, data, sign)
	if err != nil {
		t.Errorf("err: %+v\n", err)
		return
	}
	t.Logf("verify: %+v\n", verify)
	if !verify {
		t.Errorf("签名校验失败\n")
		return
	}
}
