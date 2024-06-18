// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"database/sql"
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

// Injectors from container.go:

// Dependency Injection for User Use Case
func NewUserContainer(config *commons.Config, db *sql.DB, cache2 cache.Cache, idGenerator generator.IdGenerator, fileProcessing file_statics.FileProcessing, fileUpload file_statics.FileUpload, email emails.EmailService, validator *services.Validation) *use_case.UserUseCase {
	userRepository := repository.NewUserRepositoryPG(db, idGenerator)
	passwordHash := security.NewArgon2()
	validateUser := validation.NewValidateUser(validator)
	token := security.NewJwtToken(idGenerator)
	userUseCase := use_case.NewUserUseCase(userRepository, fileProcessing, fileUpload, passwordHash, email, validateUser, config, token, cache2)
	return userUseCase
}

// Dependency Injection for Asset Use Case
func NewAssetContainer(idGenerator generator.IdGenerator, db *sql.DB, fileProcessing file_statics.FileProcessing, fileUpload file_statics.FileUpload, validator *services.Validation) *use_case.AssetUseCase {
	assetRepository := repository.NewAssetRepositoryPG(idGenerator, db)
	validateAsset := validation.NewGoValidateAsset(validator)
	assetUseCase := use_case.NewAssetUseCase(assetRepository, fileProcessing, fileUpload, validateAsset)
	return assetUseCase
}

// Dependency Injection for Favorite Use Case
func NewFavoriteContainer(idGenerator generator.IdGenerator, db *sql.DB) *use_case.FavoriteUseCase {
	favoriteRepository := repository.NewFavoriteRepositoryPG(idGenerator, db)
	favoriteUseCase := use_case.NewFavoriteUseCase(favoriteRepository)
	return favoriteUseCase
}

// Dependency Injection for Cart Use Case
func NewCartContainer(idGenerator generator.IdGenerator, db *sql.DB) *use_case.CartUseCase {
	cartRepository := repository.NewCartRepositoryPG(idGenerator, db)
	cartUseCase := use_case.NewCartUseCase(cartRepository)
	return cartUseCase
}

// Dependency Injection for Transaction Use Case
func NewTransactionContainer(idGenerator generator.IdGenerator, db *sql.DB) *use_case.TransactionUseCase {
	transactionRepository := repository.NewTransactionRepositoryPG(idGenerator, db)
	transactionUseCase := use_case.NewTransactionUseCase(transactionRepository)
	return transactionUseCase
}
