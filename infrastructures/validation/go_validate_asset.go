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

func (g *GoValidateAsset) ValidatePayload(payload *entity.AddAssetPayload) {
	schema := map[string]string{
		"title": "required,min=5,max=100",
		//"description": "max=255",
		//"details":     "max=500",
	}

	services.Validate(payload, schema, g.v)
}
