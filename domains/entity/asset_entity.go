package entity

import (
	"mime/multipart"
	"time"
)

type AddAssetPayload struct {
	Title         string `json:"title"`
	File          *multipart.FileHeader
	Description   string `json:"description"`
	Details       string `json:"details"`
	OwnerId       string
	OriginalLink  string
	WatermarkLink string
}

type PreviewAsset struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	Title       string `json:"title"`
	FilePath    string `json:"file_watermark_path"`
	Description string `json:"description"`
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
