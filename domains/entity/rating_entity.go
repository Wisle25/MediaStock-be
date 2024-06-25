package entity

import "time"

// CreateRatingPayload is the payload for creating a rating.
type CreateRatingPayload struct {
	AssetId     string `json:"assetId"`     // ID of the asset being rated
	UserId      string `json:"userId"`      // ID of the user providing the rating
	Score       int16  `json:"score"`       // Score given to the asset
	Description string `json:"description"` // Description of the rating
}

// Rating represents the rating entity.
type Rating struct {
	ID          string    `json:"id"`          // Unique identifier for the rating
	AssetID     string    `json:"assetId"`     // ID of the asset being rated
	Username    string    `json:"username"`    // Username of the user providing the rating
	UserAvatar  string    `json:"userAvatar"`  // Avatar link of the user providing the rating
	Score       int16     `json:"score"`       // Score given to the asset, SMALLINT in PostgreSQL maps to int16 in Go
	Description string    `json:"description"` // Description of the rating
	CreatedAt   time.Time `json:"createdAt"`   // Timestamp when the rating was created
}
