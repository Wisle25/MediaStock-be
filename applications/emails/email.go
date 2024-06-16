package emails

import "github.com/wisle25/media-stock-be/domains/entity"

type EmailService interface {
	SendEmail(payload entity.Email)
}
