package entity

import (
	"mime/multipart"
	"time"
)

type AssetPayload struct {
	// From User
	Title       string `json:"title"`
	File        *multipart.FileHeader
	Description string `json:"description"`
	Details     string `json:"details"`
	Category    string `json:"category"`
	Price       string `json:"price"`

	// From Server
	OwnerId       string
	OriginalLink  string
	WatermarkLink string
}

type PreviewAsset struct {
	Id            string  `json:"id"`
	OwnerUsername string  `json:"ownerUsername"`
	Title         string  `json:"title"`
	FilePath      string  `json:"filePath"`
	Category      string  `json:"category"`
	Description   string  `json:"description"`
	Rating        float32 `json:"rating"`
	FavoriteCount int     `json:"favoriteCount"`
	IsFavorite    bool    `json:"isFavorite"`
}

type Asset struct {
	Id             string    `json:"id"`
	OwnerId        string    `json:"ownerId"`
	OwnerUsername  string    `json:"ownerUsername"`
	Title          string    `json:"title"`
	FilePath       string    `json:"filePath"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	Details        string    `json:"details"`
	Rating         float32   `json:"rating"`
	Price          string    `json:"price"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	FavoriteCount  int       `json:"favoriteCount"`
	PurchasedCount int       `json:"purchasedCount"`
	IsFavorite     bool      `json:"isFavorite"`
	IsPurchased    bool      `json:"isPurchased"`
}

// SimpleAsset only provides some information
type SimpleAsset struct {
	Id       string `json:"id"`       // Unique identifier for the favorite item
	Title    string `json:"title"`    // Title of the favorited asset
	Price    string `json:"price"`    // Price of the favorited asset
	FilePath string `json:"filePath"` // Path to the favorited asset file
}
