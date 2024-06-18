package validation

import (
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/infrastructures/services"
)

type GoValidateRating struct {
	v *services.Validation
}

func NewGoValidateRating(validation *services.Validation) validation.ValidateRating {
	return &GoValidateRating{
		validation,
	}
}

func (g *GoValidateRating) ValidatePayload(payload *entity.CreateRatingPayload) {
	schema := map[string]string{
		"AssetId":     "required,uuid",
		"Score":       "required,min=1,max=5",
		"Description": "max=255",
	}

	services.Validate(payload, schema, g.v)
}
