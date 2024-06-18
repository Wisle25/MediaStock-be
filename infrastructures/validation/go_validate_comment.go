package validation

import (
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/infrastructures/services"
)

type GoValidateComment struct {
	v *services.Validation
}

func NewGoValidateComment(validation *services.Validation) validation.ValidateComment {
	return &GoValidateComment{
		validation,
	}
}

func (g *GoValidateComment) ValidatePayload(payload *entity.CreateCommentPayload) {
	schema := map[string]string{
		"AssetId": "required,uuid",
		"UserId":  "required,uuid",
		"Content": "required,max=100",
	}

	services.Validate(payload, schema, g.v)
}

func (g *GoValidateComment) ValidateUpdate(payload *entity.EditCommentPayload) {
	schema := map[string]string{
		"Content": "required,max=100",
	}

	services.Validate(payload, schema, g.v)
}
