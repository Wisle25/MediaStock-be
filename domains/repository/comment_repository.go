package repository

import "github.com/wisle25/media-stock-be/domains/entity"

// CommentRepository defines the interface for the comment repository
type CommentRepository interface {
	// CreateComment creates a new comment
	// Returning new comment id
	CreateComment(payload *entity.CreateCommentPayload) string

	// GetCommentsByAsset retrieves all comments for a specific asset
	GetCommentsByAsset(assetId string) []entity.Comment

	// UpdateComment updates an existing comment
	UpdateComment(id string, content string)

	// DeleteComment deletes a comment by its ID
	DeleteComment(id string)
}
