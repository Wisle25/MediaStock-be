package use_case

import (
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

// CommentUseCase handles the business logic for comment operations.
type CommentUseCase struct {
	commentRepository repository.CommentRepository
	validation        validation.ValidateComment
}

// NewCommentUseCase creates a new instance of CommentUseCase.
func NewCommentUseCase(
	commentRepository repository.CommentRepository,
	validation validation.ValidateComment,
) *CommentUseCase {
	return &CommentUseCase{
		commentRepository,
		validation,
	}
}

// ExecuteCreate validates the payload and creates a new comment.
func (uc *CommentUseCase) ExecuteCreate(payload *entity.CreateCommentPayload) string {
	uc.validation.ValidatePayload(payload)
	return uc.commentRepository.CreateComment(payload)
}

// ExecuteGetByAsset retrieves all comments for a specific asset.
func (uc *CommentUseCase) ExecuteGetByAsset(assetId string) []entity.Comment {
	return uc.commentRepository.GetCommentsByAsset(assetId)
}

// ExecuteUpdate validates the payload and updates an existing comment.
func (uc *CommentUseCase) ExecuteUpdate(id string, payload *entity.EditCommentPayload) {
	uc.validation.ValidateUpdate(payload)
	uc.commentRepository.UpdateComment(id, payload.Content)
}

// ExecuteDelete deletes a comment by its ID.
func (uc *CommentUseCase) ExecuteDelete(id string) {
	uc.commentRepository.DeleteComment(id)
}
