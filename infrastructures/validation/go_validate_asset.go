package validation

import (
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/infrastructures/services"
)

type GoValidateAsset struct {
	v *services.Validation
}

func NewGoValidateAsset(validation *services.Validation) validation.ValidateAsset {
	return &GoValidateAsset{
		validation,
	}
}

func (g *GoValidateAsset) ValidatePayload(payload *entity.AssetPayload) {
	schema := map[string]string{
		"Title":       "required,min=5,max=100",
		"Price":       "required",
		"Description": "required,max=255",
		"Category":    "required,oneof='Panorama' 'City and Architecture' 'Peoples and Portraits' 'Foods and Drink' 'Animals' 'Object'",
		"Details":     "max=500",
	}

	services.Validate(payload, schema, g.v)
}
