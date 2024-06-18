package validation

import "github.com/wisle25/media-stock-be/domains/entity"

type ValidateRating interface {
	ValidatePayload(payload *entity.CreateRatingPayload)
}
