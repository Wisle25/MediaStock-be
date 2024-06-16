package repository_test

import (
	"github.com/wisle25/media-stock-be/commons"
	"github.com/wisle25/media-stock-be/infrastructures/generator"
	"github.com/wisle25/media-stock-be/infrastructures/repository"
	"github.com/wisle25/media-stock-be/infrastructures/services"
	"github.com/wisle25/media-stock-be/tests/db_helper"
	"testing"
)

func TestAssetRepositoryPG(t *testing.T) {
	// Arrange
	config := commons.LoadConfig("../..")
	db := services.ConnectDB(config)
	userHelperDb := &db_helper.UserHelperDB{
		DB: db,
	}
	defer userHelperDb.CleanUserDB()

	uuidGenerator := generator.NewUUIDGenerator()
	assetRepositoryPG := repository.NewAssetRepositoryPG(db, uuidGenerator)

	t.Run("Add Repository", func(t *testing.T) {
		
	})
}
