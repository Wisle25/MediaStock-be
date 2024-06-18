package validation

import "github.com/wisle25/media-stock-be/domains/entity"

type ValidateComment interface {
	ValidatePayload(payload *entity.CreateCommentPayload)
	ValidateUpdate(payload *entity.EditCommentPayload)
}
