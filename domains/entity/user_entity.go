package entity

import "mime/multipart"

// RegisterUserPayload represents the payload for user registration.
type RegisterUserPayload struct {
	Username        string `json:"username"`        // Username chosen by the user
	Password        string `json:"password"`        // Password chosen by the user
	Email           string `json:"email"`           // User's email address
	ConfirmPassword string `json:"confirmPassword"` // Confirmation of the user's password
}

// LoginUserPayload represents the payload for user login.
type LoginUserPayload struct {
	Identity string `json:"identity"` // User's identity which could be username or email
	Password string `json:"password"` // User's password
}

// UpdateUserPayload represents the payload for updating user information.
type UpdateUserPayload struct {
	Username        string                `json:"username"`        // New username for the user
	Email           string                `json:"email"`           // New email address for the user
	Password        string                `json:"password"`        // New password for the user
	ConfirmPassword string                `json:"confirmPassword"` // Confirmation of the new password
	Avatar          *multipart.FileHeader // New avatar image for the user
}

// User represents a user in the system.
type User struct {
	Id         string `json:"id"`         // ID of the user
	Username   string `json:"username"`   // Username of the user, should be unique
	Email      string `json:"email"`      // Email address of the user, should be unique
	AvatarLink string `json:"avatarLink"` // Link to the user's avatar image
	IsVerified bool   `json:"isVerified"` // Indicates whether the user's email is verified
}
