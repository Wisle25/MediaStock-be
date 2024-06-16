package entity

import (
	"mime/multipart"
	"time"
)

type AddAssetPayload struct {
	Title         string                `json:"title"`
	File          *multipart.FileHeader `json:"-"`
	Description   string                `json:"description"`
	Details       string                `json:"details"`
	OwnerId       string
	OriginalLink  string
	WatermarkLink string
}

type PreviewAsset struct {
	Id          string
	OwnerId     string
	Title       string
	FilePath    string
	Description string
}

type Asset struct {
	Id          string
	OwnerId     string
	Title       string
	FilePath    string
	Description string
	Details     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
