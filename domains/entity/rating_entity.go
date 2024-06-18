package entity

import "time"

// CreateRatingPayload is the payload for creating a rating
type CreateRatingPayload struct {
	AssetId     string `json:"asset_id"`
	UserId      string
	Score       int16  `json:"score"`
	Description string `json:"description"`
}

// Rating represents the rating entity
type Rating struct {
	ID          string    `json:"id"`
	AssetID     string    `json:"asset_id"`
	Username    string    `json:"username"`
	UserAvatar  string    `json:"user_avatar"`
	Score       int16     `json:"score"` // SMALLINT in PostgreSQL maps to int16 in Go
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
