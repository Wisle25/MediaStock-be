package entity

// Email represents the data required to send an email.
type Email struct {
	To      string // Recipient email address
	Subject string // Subject of the email
	Body    string // Body of the email
}
