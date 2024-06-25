package use_case

import (
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

// RatingUseCase handles the business logic for rating operations.
type RatingUseCase struct {
	ratingRepository repository.RatingRepository
	validation       validation.ValidateRating
}

// NewRatingUseCase creates a new instance of RatingUseCase.
func NewRatingUseCase(
	ratingRepository repository.RatingRepository,
	validation validation.ValidateRating,
) *RatingUseCase {
	return &RatingUseCase{
		ratingRepository,
		validation,
	}
}

// ExecuteCreate validates the payload and creates a new rating.
func (uc *RatingUseCase) ExecuteCreate(payload *entity.CreateRatingPayload) string {
	uc.validation.ValidatePayload(payload)
	return uc.ratingRepository.CreateRating(payload)
}

// ExecuteGetByAsset retrieves all ratings for a specific asset.
func (uc *RatingUseCase) ExecuteGetByAsset(assetId string) []entity.Rating {
	return uc.ratingRepository.GetRatingsByAsset(assetId)
}

// ExecuteDelete deletes a rating by its ID.
func (uc *RatingUseCase) ExecuteDelete(id string) {
	uc.ratingRepository.DeleteRating(id)
}
