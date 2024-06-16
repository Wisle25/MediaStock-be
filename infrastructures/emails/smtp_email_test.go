package emails_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/infrastructures/emails"
	"testing"
)

func TestSendEmail(t *testing.T) {
	// Arrange
	config := commons.LoadConfig("../../")
	emailService := emails.NewStmpEmailService(config)
	email := entity.Email{
		To:      "handidwic1225@gmail.com",
		Subject: "Hello World!",
		Body:    "Hello World! This is test from my app",
	}

	// Action and Assert
	assert.NotPanics(t, func() {
		emailService.SendEmail(email)
	})
}
