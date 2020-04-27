package models

import "github.com/jinzhu/gorm"

// Asset model
type Asset struct {
	gorm.Model
	UserID   int    `gorm:"not null"`
	AppID    int    `gorm:"not null"`
	Type     string `gorm:"not null"`
	Attrs    Attrs
	StoreRef string `json:"-"`
}

// Attrs for unstructured JSON associated with an asset
type Attrs map[string]interface{}
