package model

import "gorm.io/gorm"

type Data struct {
	DataType string `gorm:"data_type" json:"data_type"`
	Data     string `gorm:"data" json:"-"`
	DataHash string `gorm:"data_hash" json:"data_hash"`
}

type DataModel struct {
	gorm.Model
	Data
}

func (DataModel) TableName() string {
	return "data"
}

type DataInquiry struct {
	DataModel
	DataHashes []string `json:"data_hashes"`
}
