package models

import "github.com/jinzhu/gorm"

// Asset model
type Asset struct {
	gorm.Model
	UserID   int    `gorm:"not null"`
	AppID    int    `gorm:"not null"`
	Type     string `gorm:"not null"`
	Attrs    string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
	StoreRef string `json:"-"`
}
