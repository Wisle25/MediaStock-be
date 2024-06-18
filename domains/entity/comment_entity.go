package entity

import "time"

type CreateCommentPayload struct {
	AssetId string `json:"asset_id"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
}

type EditCommentPayload struct {
	Content string `json:"content"`
}

// Comment represents the comment entity
type Comment struct {
	Id         string    `json:"id"`
	AssetId    string    `json:"asset_id"`
	UserId     string    `json:"user_id"`
	Username   string    `json:"username"`
	AvatarLink string    `json:"avatar_link"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
