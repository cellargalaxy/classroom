package model

import (
	"time"
)

type User struct {
	Id             int32     `gorm:"id" json:"id"`
	PublicKey      string    `gorm:"public_key" json:"-"`
	PublicKeyHash  string    `gorm:"public_key_hash" json:"public_key_hash"`
	PrivateKey     string    `gorm:"private_key" json:"-"`
	PrivateKeyHash string    `gorm:"private_key_hash" json:"private_key_hash"`
	CreatedAt      time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"updated_at" json:"updated_at"`
}

type UserAdd struct {
	User
	Sign string `json:"sign"`
}

type Message struct {
	Id                    int32     `gorm:"id" json:"id"`
	DataType              string    `gorm:"data_type" json:"data_type"`
	Data                  string    `gorm:"data" json:"-"`
	DataHash              string    `gorm:"data_hash" json:"data_hash"`
	CreateSign            string    `gorm:"create_sign" json:"-"`
	CreatePublicKeyHash   string    `gorm:"create_public_key_hash" json:"create_public_key_hash"`
	PublishSign           string    `gorm:"publish_sign" json:"-"`
	PublishPrivateKeyHash string    `gorm:"publish_private_key_hash" json:"publish_private_key_hash"`
	PublishSignHash       string    `gorm:"publish_sign_hash" json:"publish_sign_hash"`
	ParentHash            string    `gorm:"parent_hash" json:"parent_hash"`
	PublishAt             time.Time `gorm:"publish_at" json:"publish_at"`
	CreatedAt             time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt             time.Time `gorm:"updated_at" json:"updated_at"`
}

type MessageAdd struct {
	Message
}
