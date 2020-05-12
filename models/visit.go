package models

import (
	"time"
)

// Visit model
type Visit struct {
	AppID     int
	UserID    int
	Type      string
	IP        string
	UserAgent string
	APIKey    string
	VisitedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"visited_at"`
}
