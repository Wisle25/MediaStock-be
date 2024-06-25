package repository

import "github.com/wisle25/media-stock-be/domains/entity"

// CommentRepository defines the methods that any data storage provider needs to implement to manage comment data.
type CommentRepository interface {
	// CreateComment creates a new comment.
	// Returns the ID of the newly created comment.
	CreateComment(payload *entity.CreateCommentPayload) string

	// GetCommentsByAsset retrieves all comments for a specific asset by its ID.
	GetCommentsByAsset(assetId string) []entity.Comment

	// UpdateComment updates the content of an existing comment by its ID.
	UpdateComment(id string, content string)

	// DeleteComment deletes a comment by its ID.
	DeleteComment(id string)
}
