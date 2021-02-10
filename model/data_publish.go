package model

import "gorm.io/gorm"

type DataPublish struct {
	ParentHash      string `gorm:"parent_hash" json:"parent_hash"`
	CreateSignHash  string `gorm:"create_sign_hash" json:"create_sign_hash"`
	PublicKeyHash   string `gorm:"public_key_hash" json:"public_key_hash"`
	PublishSign     string `gorm:"publish_sign" json:"-"`
	PublishSignHash string `gorm:"publish_sign_hash" json:"publish_sign_hash"`
}

type DataPublishModel struct {
	gorm.Model
	DataPublish
}

func (DataPublishModel) TableName() string {
	return "data_publish"
}
