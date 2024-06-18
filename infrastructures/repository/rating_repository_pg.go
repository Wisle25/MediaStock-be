package repository

import (
	"database/sql"
	"fmt"
	"github.com/wisle25/media-stock-be/infrastructures/services"

	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/generator"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
)

type RatingRepositoryPG struct {
	db          *sql.DB
	idGenerator generator.IdGenerator
}

func NewRatingRepositoryPG(db *sql.DB, idGenerator generator.IdGenerator) repository.RatingRepository {
	return &RatingRepositoryPG{
		db:          db,
		idGenerator: idGenerator,
	}
}

func (r *RatingRepositoryPG) CreateRating(payload *entity.CreateRatingPayload) string {
	// Create ID
	id := r.idGenerator.Generate()

	// Query
	query := `
		INSERT INTO 
		    ratings(id, asset_id, user_id, score, description) 
		VALUES 
		    ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var returnedId string
	err := r.db.QueryRow(
		query,
		id,
		payload.AssetId,
		payload.UserId,
		payload.Score,
		payload.Description,
	).Scan(&returnedId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("rating_repo_pg_error: create rating: %v", err))
	}

	return returnedId
}

func (r *RatingRepositoryPG) GetRatingsByAsset(assetId string) []entity.Rating {
	// Query
	query := `
		SELECT 
		    r.id, 
		    r.asset_id, 
		    u.username AS username, 
		    u.avatar_link AS user_avatar, 
		    r.score, 
		    r.description, 
		    r.created_at 
		FROM ratings r
		INNER JOIN users u ON u.id = r.user_id
		WHERE asset_id = $1
	`

	rows, err := r.db.Query(query, assetId)
	if err != nil {
		panic(fmt.Errorf("rating_repo_pg_error: get ratings by asset: %v", err))
	}

	return services.GetTableDB[entity.Rating](rows)
}

func (r *RatingRepositoryPG) DeleteRating(id string) {
	// Query
	query := `DELETE FROM ratings WHERE id = $1`
	result, err := r.db.Exec(query, id)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("rating_repo_pg_error: delete rating: %v", err))
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		panic(fiber.NewError(fiber.StatusNotFound, "Rating not found!"))
	}
}
