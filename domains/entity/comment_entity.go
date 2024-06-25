package entity

import "time"

// CreateCommentPayload represents the data required to create a new comment.
type CreateCommentPayload struct {
	AssetId string `json:"assetId"` // ID of the asset the comment is associated with
	UserId  string `json:"userId"`  // ID of the user creating the comment
	Content string `json:"content"` // Content of the comment
}

// EditCommentPayload represents the data required to edit an existing comment.
type EditCommentPayload struct {
	Content string `json:"content"` // Updated content of the comment
}

// Comment represents the comment entity.
type Comment struct {
	Id         string    `json:"id"`         // Unique identifier for the comment
	AssetId    string    `json:"assetId"`    // ID of the asset the comment is associated with
	UserId     string    `json:"userId"`     // ID of the user who created the comment
	Username   string    `json:"username"`   // Username of the user who created the comment
	AvatarLink string    `json:"avatarLink"` // Link to the avatar of the user who created the comment
	Content    string    `json:"content"`    // Content of the comment
	CreatedAt  time.Time `json:"createdAt"`  // Timestamp when the comment was created
	UpdatedAt  time.Time `json:"updatedAt"`  // Timestamp when the comment was last updated
}
