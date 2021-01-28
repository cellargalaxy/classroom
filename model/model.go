package model

import (
	"time"
)

func (User) TableName() string {
	return "user"
}

type User struct {
	Id            int32     `gorm:"id" json:"id"`
	PublicKey     string    `gorm:"public_key" json:"-"`
	PublicKeyHash string    `gorm:"public_key_hash" json:"public_key_hash"`
	PrivateKey    string    `gorm:"private_key" json:"-"`
	CreatedAt     time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at" json:"updated_at"`
}

type UserAdd struct {
	User
	Sign string `json:"sign"`
}

func (Message) TableName() string {
	return "message"
}

type Message struct {
	Id                   int32     `gorm:"id" json:"id"`
	DataType             string    `gorm:"data_type" json:"data_type"`
	Data                 string    `gorm:"data" json:"-"`
	DataHash             string    `gorm:"data_hash" json:"data_hash"`
	CreateSign           string    `gorm:"create_sign" json:"-"`
	CreatePublicKeyHash  string    `gorm:"create_public_key_hash" json:"create_public_key_hash"`
	CreateSignHash       string    `gorm:"create_sign_hash" json:"create_sign_hash"`
	PublishSign          string    `gorm:"publish_sign" json:"-"`
	PublishPublicKeyHash string    `gorm:"publish_public_key_hash" json:"publish_public_key_hash"`
	PublishSignHash      string    `gorm:"publish_sign_hash" json:"publish_sign_hash"`
	ParentHash           string    `gorm:"parent_hash" json:"parent_hash"`
	PublishAt            time.Time `gorm:"publish_at" json:"publish_at"`
	CreatedAt            time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt            time.Time `gorm:"updated_at" json:"updated_at"`
}

type MessageAdd struct {
	Message
}

type MessageInquiry struct {
	Message
	PublishAtStart time.Time `json:"publish_at_start"`
	PublishAtEnd   time.Time `json:"publish_at_end"`
	Offset         int       `json:"offset"`
	Length         int       `json:"length"`
}
