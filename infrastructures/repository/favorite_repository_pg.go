package repository

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/generator"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
	"github.com/wisle25/media-stock-be/infrastructures/services"
)

type FavoriteRepositoryPG struct {
	idGenerator generator.IdGenerator
	db          *sql.DB
}

func NewFavoriteRepositoryPG(idGenerator generator.IdGenerator, db *sql.DB) repository.FavoriteRepository {
	return &FavoriteRepositoryPG{
		idGenerator,
		db,
	}
}

func (f *FavoriteRepositoryPG) AddAsFavorite(payload *entity.FavoritePayload) {
	id := f.idGenerator.Generate()

	// Query
	query := "INSERT INTO favorites(id, asset_id, user_id) VALUES ($1, $2, $3)"
	result, err := f.db.Exec(query, id, payload.AssetId, payload.UserId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("add_favorite_err: %v", err))
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		panic(fiber.NewError(fiber.StatusBadRequest, "Unable to add favorite"))
	}
}

func (f *FavoriteRepositoryPG) GetAllFavorites(userId string) []entity.Favorite {
	// Query
	query := `
			SELECT 
    			a.id,
				a.title,
				a.price,
			    a.file_watermark_path AS file_path
			FROM favorites f
			INNER JOIN assets a ON f.asset_id = a.id
			WHERE user_id = $1`
	rows, err := f.db.Query(query, userId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("get_all_favorites_err: %v", err))
	}

	return services.GetTableDB[entity.Favorite](rows)
}

func (f *FavoriteRepositoryPG) RemoveFavorite(payload *entity.FavoritePayload) {
	// Query
	query := "DELETE FROM favorites WHERE asset_id = $1 AND user_id = $2"
	result, err := f.db.Exec(query, payload.AssetId, payload.UserId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("remove_favorite_err: %v", err))
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		panic(fiber.NewError(fiber.StatusBadRequest, "Unable to remove favorite"))
	}
}
