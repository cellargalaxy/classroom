package model

import (
	"gorm.io/gorm"
	"time"
)

type Inquiry struct {
	StartCreatedAt time.Time
	EndCreatedAt   time.Time
	Offset         int
	Limit          int
}

type User struct {
	gorm.Model
	PublicKey     string
	PublicKeyHash string
	PrivateKey    string
}

type UserAdd struct {
	User
	Sign string
}

type Message struct {
	gorm.Model
	UserHash    string
	ParentHash  string
	DataType    string
	Data        string
	Hash        string
	SummaryType string
	Summary     string
}

type MessageInquiry struct {
	Message
	Inquiry
}
