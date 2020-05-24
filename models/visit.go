package models

import (
	"time"
)

// Visit model
type Visit struct {
	ObjectType string    `gorm:"not null" json:"obj_type"`
	ObjectID   int       `gorm:"not null" json:"obj_id"`
	Type       string    `json:"type"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
	APIKey     string    `gorm:"not null" json:"-"`
	VisitedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"visited_at"`
}
