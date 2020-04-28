package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// Attrs is for json
type Attrs interface{}

// Asset model
type Asset struct {
	gorm.Model
	AppID    int    `gorm:"index;not null"`
	UserID   int    `gorm:"not null"`
	FileName string `gorm:"not null"`
	Type     string `gorm:"not null"`
	Attrs    postgres.Jsonb
}
