//go:build wireinject
// +build wireinject

package container

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/wisle25/media-stock-be/applications/cache"
	"github.com/wisle25/media-stock-be/applications/emails"
	"github.com/wisle25/media-stock-be/applications/file_statics"
	"github.com/wisle25/media-stock-be/applications/generator"
	"github.com/wisle25/media-stock-be/applications/use_case"
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/infrastructures/repository"
	"github.com/wisle25/media-stock-be/infrastructures/security"
	"github.com/wisle25/media-stock-be/infrastructures/services"
	"github.com/wisle25/media-stock-be/infrastructures/validation"
)

// Dependency Injection for User Use Case
func NewUserContainer(
	config *commons.Config,
	db *sql.DB,
	cache cache.Cache,
	idGenerator generator.IdGenerator,
	fileProcessing file_statics.FileProcessing,
	fileUpload file_statics.FileUpload,
	email emails.EmailService,
	validator *services.Validation,
) *use_case.UserUseCase {
	wire.Build(
		repository.NewUserRepositoryPG,
		security.NewArgon2,
		validation.NewValidateUser,
		security.NewJwtToken,
		use_case.NewUserUseCase,
	)

	return nil
}

// Dependency Injection for Asset Use Case
func NewAssetContainer(
	idGenerator generator.IdGenerator,
	db *sql.DB,
	fileProcessing file_statics.FileProcessing,
	fileUpload file_statics.FileUpload,
	validator *services.Validation,
) *use_case.AssetUseCase {
	wire.Build(
		validation.NewGoValidateAsset,
		repository.NewAssetRepositoryPG,
		use_case.NewAssetUseCase,
	)

	return nil
}

// Dependency Injection for Favorite Use Case
func NewFavoriteContainer(
	idGenerator generator.IdGenerator,
	db *sql.DB,
) *use_case.FavoriteUseCase {
	wire.Build(
		repository.NewFavoriteRepositoryPG,
		use_case.NewFavoriteUseCase,
	)

	return nil
}

// Dependency Injection for Cart Use Case
func NewCartContainer(
	idGenerator generator.IdGenerator,
	db *sql.DB,
) *use_case.CartUseCase {
	wire.Build(
		repository.NewCartRepositoryPG,
		use_case.NewCartUseCase,
	)

	return nil
}
