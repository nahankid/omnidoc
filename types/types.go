package types

import (
	"encoding/json"
	"omnidoc/lib"
	"omnidoc/models"
)

// CreateRequest struct
type CreateRequest struct {
	AppID    int             `json:"app_id"`
	UserID   int             `json:"user_id"`
	FileName string          `json:"filename"`
	Type     string          `json:"type"`
	Attrs    json.RawMessage `json:"attrs"`
}

// GetResponse struct
type GetResponse struct {
	Asset     models.Asset       `json:"asset"`
	SignedURL lib.S3PresignedURL `json:"signed_url"`
}
