package model

import "gorm.io/gorm"

type User struct {
	PublicKey      string `gorm:"public_key" json:"-"`
	PublicKeyHash  string `gorm:"public_key_hash" json:"public_key_hash"`
	PrivateKey     string `gorm:"private_key" json:"-"`
	PrivateKeyHash string `gorm:"private_key_hash" json:"private_key_hash"`
}

type UserModel struct {
	gorm.Model
	User
}

func (UserModel) TableName() string {
	return "user"
}

type UserAdd struct {
	User
	Sign string `json:"sign"`
}
