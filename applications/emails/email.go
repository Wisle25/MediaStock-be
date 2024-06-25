package emails

import "github.com/wisle25/media-stock-be/domains/entity"

// EmailService defines the interface for an email service.
// Any struct that implements this interface can be used to send emails.
type EmailService interface {
	// SendEmail sends an email based on the provided payload.
	// The payload should contain necessary email details like recipient, subject, and body.
	SendEmail(payload entity.Email)
}
