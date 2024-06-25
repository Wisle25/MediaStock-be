package validation

import "github.com/wisle25/media-stock-be/domains/entity"

type ValidateAsset interface {
	ValidatePayload(payload *entity.AssetPayload)
}
