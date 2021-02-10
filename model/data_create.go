package model

import "gorm.io/gorm"

type DataCreate struct {
	DataHash       string `gorm:"data_hash" json:"data_hash"`
	PublicKeyHash  string `gorm:"public_key_hash" json:"public_key_hash"`
	CreateSign     string `gorm:"create_sign" json:"-"`
	CreateSignHash string `gorm:"create_sign_hash" json:"create_sign_hash"`
}

type DataCreateModel struct {
	gorm.Model
	DataCreate
}

func (DataCreateModel) TableName() string {
	return "data_create"
}

type DataCreateAdd struct {
	Data
	DataCreate
}

type DataCreateAddModel struct {
	DataModel
	DataCreateModel
}
