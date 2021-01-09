package model

import (
	"gorm.io/gorm"
	"time"
)

type Inquiry struct {
	StartCreatedAt time.Time `json:"start_created_at"`
	EndCreatedAt   time.Time `json:"end_created_at"`
	Offset         int       `json:"offset"`
	Limit          int       `json:"limit"`
}

type User struct {
	gorm.Model
	PublicKey     string `gorm:"public_key" json:"public_key"`
	PublicKeyHash string `gorm:"public_key_hash" json:"public_key_hash"`
	PrivateKey    string `gorm:"private_key" json:"private_key"`
}

type UserAdd struct {
	User
	Sign string `json:"sign"`
}

type Message struct {
	gorm.Model
	UserHash   string `gorm:"user_hash" json:"user_hash"`
	ParentHash string `gorm:"parent_hash" json:"parent_hash"`
	Data       string `gorm:"data" json:"data"`
	DataType   string `gorm:"data_type" json:"data_type"`
	Hash       string `gorm:"hash" json:"hash"`
}

type MessageAdd struct {
	Message
	Sign string `json:"sign"`
}

type MessageInquiry struct {
	Message
	Inquiry
}
