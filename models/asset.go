package models

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// Attrs is for json
type Attrs interface{}

// Asset model
type Asset struct {
	APIKey     string         `gorm:"not null" json:"-"`
	ObjectType string         `gorm:"not null" json:"obj_type"`
	ObjectID   int            `gorm:"index;not null" json:"obj_id"`
	FileName   string         `gorm:"not null" json:"filename"`
	Type       string         `gorm:"not null" json:"type"`
	Attrs      postgres.Jsonb `json:"attrs"`
	CreatedAt  time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"not null" json:"updated_at"`
}
