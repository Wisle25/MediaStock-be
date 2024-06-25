package use_case

import (
	"encoding/json"
	"fmt"
	"github.com/wisle25/media-stock-be/applications/cache"
	"github.com/wisle25/media-stock-be/applications/emails"
	"github.com/wisle25/media-stock-be/applications/file_statics"
	"github.com/wisle25/media-stock-be/applications/security"
	"github.com/wisle25/media-stock-be/applications/validation"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
	"io"
	"path"
	"time"
)

// UserUseCase handles the business logic for user operations.
type UserUseCase struct {
	userRepository repository.UserRepository
	fileProcessing file_statics.FileProcessing
	fileUpload     file_statics.FileUpload
	passwordHash   security.PasswordHash
	emailService   emails.EmailService
	validator      validation.ValidateUser
	config         *commons.Config
	token          security.Token
	cache          cache.Cache
}

func NewUserUseCase(
	userRepository repository.UserRepository,
	fileProcessing file_statics.FileProcessing,
	fileUpload file_statics.FileUpload,
	passwordHash security.PasswordHash,
	emailService emails.EmailService,
	validator validation.ValidateUser,
	config *commons.Config,
	token security.Token,
	cache cache.Cache,
) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
		fileProcessing: fileProcessing,
		fileUpload:     fileUpload,
		passwordHash:   passwordHash,
		emailService:   emailService,
		validator:      validator,
		config:         config,
		token:          token,
		cache:          cache,
	}
}

// ExecuteAdd Handling user registration.
// Should raise panic if violates username/email uniqueness.
// Returning registered user's ID.
func (uc *UserUseCase) ExecuteAdd(payload *entity.RegisterUserPayload) string {
	uc.validator.ValidateRegisterPayload(payload)

	payload.Password = uc.passwordHash.Hash(payload.Password)

	registeredId := uc.userRepository.AddUser(payload)

	go uc.SendEmailActivation(payload, registeredId)

	return registeredId
}

// SendEmailActivation sending email activation to user
func (uc *UserUseCase) SendEmailActivation(payload *entity.RegisterUserPayload, registeredId string) {
	activationToken := uc.token.CreateToken(
		&entity.User{
			Id: registeredId,
		},
		uc.config.AccessTokenExpiresIn,
		uc.config.AccessTokenPrivateKey,
	)
	activationURL := fmt.Sprintf(
		"%s://%s:%s/activate?token=%s",
		uc.config.ServerProtocol,
		uc.config.ServerHost,
		uc.config.ServerPort,
		activationToken.Token,
	)
	activationEmail := entity.Email{
		To:      payload.Email,
		Subject: "Account Activation",
		Body: fmt.Sprintf(`
			<h2>Welcome %s!</h2>
			<p>Thank you for registering. Please click the button below to activate your account:</p>
			<a href="%s" style="display: inline-block; padding: 10px 20px; font-size: 16px; color: white; background-color: #007BFF; text-decoration: none; border-radius: 5px;">Activate Here</a>
			<p>If the button above does not work, copy and paste the following link into your browser:</p>
			<p><a href="%s">%s</a></p>
		`, payload.Username, activationURL, activationURL, activationURL),
	}
	uc.emailService.SendEmail(activationEmail)
}

// ExecuteActivate Activating user from email (email verification)
// Receiving payload token, so it can get user id when validating it
func (uc *UserUseCase) ExecuteActivate(payloadToken string) {
	tokenDetail := uc.token.ValidateToken(payloadToken, uc.config.AccessTokenPublicKey)
	uc.userRepository.ActivateAccount(tokenDetail.UserToken.Id)
}

// ExecuteLogin Handling user login. Returning user's token for authentication/authorization later.
// Should raise panic if user is not existed
// Returned tokens must be added to the HTTP cookie
func (uc *UserUseCase) ExecuteLogin(payload *entity.LoginUserPayload) (*entity.TokenDetail, *entity.TokenDetail) {
	uc.validator.ValidateLoginPayload(payload)

	// Get user information from database then compare password
	userInfo, encryptedPassword := uc.userRepository.GetUserForLogin(payload.Identity)
	uc.passwordHash.Compare(payload.Password, encryptedPassword)

	// Create token
	accessTokenDetail := uc.token.CreateToken(userInfo, uc.config.AccessTokenExpiresIn, uc.config.AccessTokenPrivateKey)
	refreshTokenDetail := uc.token.CreateToken(userInfo, uc.config.RefreshTokenExpiresIn, uc.config.RefreshTokenPrivateKey)

	// Add tokens to the cache
	userInfoJSON, err := json.Marshal(userInfo)
	if err != nil {
		commons.ThrowServerError("login_err: unable to marshal json user info", err)
	}

	now := time.Now()
	uc.cache.SetCache(accessTokenDetail.TokenId, userInfoJSON, time.Unix(accessTokenDetail.ExpiresIn, 0).Sub(now))
	uc.cache.SetCache(refreshTokenDetail.TokenId, userInfoJSON, time.Unix(refreshTokenDetail.ExpiresIn, 0).Sub(now))

	// Returned token should be added to HTTP Cookie
	return accessTokenDetail, refreshTokenDetail
}

// ExecuteRefreshToken handles refreshing the access token using the provided refresh token.
// Should raise panic if refresh token is invalid
// Returned new access token should be added to HTTP Cookie
func (uc *UserUseCase) ExecuteRefreshToken(currentRefreshToken string) *entity.TokenDetail {
	// Verify token from JWT itself and from cache
	tokenClaims := uc.token.ValidateToken(currentRefreshToken, uc.config.RefreshTokenPublicKey)
	userInfoJSON := uc.cache.GetCache(tokenClaims.TokenId).(string)

	// Unmarshal user info JSON
	var userInfo entity.User
	err := json.Unmarshal([]byte(userInfoJSON), &userInfo)
	if err != nil {
		commons.ThrowServerError("refresh_token_err: unable to unmarshal json user info", err)
	}

	// Re-create access token and re-insert to the cache
	now := time.Now()
	accessTokenDetail := uc.token.CreateToken(&userInfo, uc.config.AccessTokenExpiresIn, uc.config.AccessTokenPrivateKey)
	uc.cache.SetCache(accessTokenDetail.TokenId, userInfoJSON, time.Unix(accessTokenDetail.ExpiresIn, 0).Sub(now))

	// Returned token should be added to HTTP Cookie
	return accessTokenDetail
}

// ExecuteLogout handles user logout by removing the tokens from the cache.
// Don't forget to remove the tokens from cookies too in infrastructure layer
func (uc *UserUseCase) ExecuteLogout(refreshToken string, accessTokenId string) {
	// Verify
	refreshTokenClaims := uc.token.ValidateToken(refreshToken, uc.config.RefreshTokenPublicKey)

	// Remove from cache
	uc.cache.DeleteCache(refreshTokenClaims.TokenId)
	uc.cache.DeleteCache(accessTokenId)
}

// ExecuteGuard verifies the access token and retrieves the associated user from the cache.
// This is used as a guard middleware for JWT authentication.
// Returning userId from token's cache
func (uc *UserUseCase) ExecuteGuard(accessToken string) (interface{}, *entity.TokenDetail) {
	accessTokenDetail := uc.token.ValidateToken(accessToken, uc.config.AccessTokenPublicKey)

	return uc.cache.GetCache(accessTokenDetail.TokenId), accessTokenDetail
}

// ExecuteGetUserById simply returns specified user information by ID
func (uc *UserUseCase) ExecuteGetUserById(userId string) *entity.User {
	return uc.userRepository.GetUserById(userId)
}

// ExecuteUpdateUserById Updating user information and now user can set their new password and upload an avatar.
func (uc *UserUseCase) ExecuteUpdateUserById(userId string, payload *entity.UpdateUserPayload) {
	uc.validator.ValidateUpdatePayload(payload)

	// Hash password if provided only
	if payload.Password != "" {
		payload.Password = uc.passwordHash.Hash(payload.Password)
	}

	newAvatarLink := ""

	// Handling avatar file
	if payload.Avatar != nil {
		file, _ := payload.Avatar.Open()
		fileBuffer, _ := io.ReadAll(file)

		compressedBuffer, extension := uc.fileProcessing.CompressImage(fileBuffer, file_statics.WEBP)
		newAvatarLink = uc.fileUpload.UploadFile(compressedBuffer, extension)

		// Adjust the link to be added with Minio
		minioUrl := uc.config.MinioUrl + uc.config.MinioBucket + "/"
		newAvatarLink = minioUrl + newAvatarLink
	}

	// Updating user's repository
	oldAvatarLink := uc.userRepository.UpdateUserById(userId, payload, newAvatarLink)

	// If exists, remove user's old avatar
	if oldAvatarLink != "" {
		uc.fileUpload.RemoveFile(path.Base(oldAvatarLink))
	}
}
