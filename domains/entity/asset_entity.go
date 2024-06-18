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
	Price         string `json:"price"`
	OwnerId       string
	OriginalLink  string
	WatermarkLink string
}

type PreviewAsset struct {
	Id            string `json:"id"`
	OwnerUsername string `json:"owner_username"`
	Title         string `json:"title"`
	FilePath      string `json:"file_watermark_path"`
	Description   string `json:"description"`
	FavoriteCount int    `json:"favorite_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

type Asset struct {
	Id             string    `json:"id"`
	OwnerId        string    `json:"owner_id"`
	OwnerUsername  string    `json:"owner_username"`
	Title          string    `json:"title"`
	FilePath       string    `json:"file_watermark_path"`
	Description    string    `json:"description"`
	Details        string    `json:"details"`
	Price          string    `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	FavoriteCount  int       `json:"favorite_count"`
	PurchasedCount int       `json:"purchased_count"`
	IsFavorite     bool      `json:"is_favorite"`
	IsPurchased    bool      `json:"is_purchased"`
}
