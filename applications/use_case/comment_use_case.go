package use_case

import (
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

type CommentUseCase struct {
	commentRepository repository.CommentRepository
	validation        validation.ValidateComment
}

func NewCommentUseCase(
	commentRepository repository.CommentRepository,
	validation validation.ValidateComment,
) *CommentUseCase {
	return &CommentUseCase{
		commentRepository,
		validation,
	}
}

func (uc *CommentUseCase) ExecuteCreate(payload *entity.CreateCommentPayload) string {
	uc.validation.ValidatePayload(payload)
	return uc.commentRepository.CreateComment(payload)
}

func (uc *CommentUseCase) ExecuteGetByAsset(assetId string) []entity.Comment {
	return uc.commentRepository.GetCommentsByAsset(assetId)
}

func (uc *CommentUseCase) ExecuteUpdate(id string, payload *entity.EditCommentPayload) {
	uc.validation.ValidateUpdate(payload)
	uc.commentRepository.UpdateComment(id, payload.Content)
}

func (uc *CommentUseCase) ExecuteDelete(id string) {
	uc.commentRepository.DeleteComment(id)
}
