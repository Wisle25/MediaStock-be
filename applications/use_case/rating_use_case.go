package use_case

import (
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

type RatingUseCase struct {
	ratingRepository repository.RatingRepository
	validation       validation.ValidateRating
}

func NewRatingUseCase(
	ratingRepository repository.RatingRepository,
	validation validation.ValidateRating,
) *RatingUseCase {
	return &RatingUseCase{
		ratingRepository,
		validation,
	}
}

func (uc *RatingUseCase) ExecuteCreate(payload *entity.CreateRatingPayload) string {
	uc.validation.ValidatePayload(payload)

	return uc.ratingRepository.CreateRating(payload)
}

func (uc *RatingUseCase) ExecuteGetByAsset(assetId string) []entity.Rating {
	return uc.ratingRepository.GetRatingsByAsset(assetId)
}

func (uc *RatingUseCase) ExecuteDelete(id string) {
	uc.ratingRepository.DeleteRating(id)
}
