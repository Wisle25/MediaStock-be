package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/generator"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
	"github.com/wisle25/media-stock-be/infrastructures/services"
)

type AssetRepositoryPG struct {
	idGenerator generator.IdGenerator
	db          *sql.DB
}

func NewAssetRepositoryPG(idGenerator generator.IdGenerator, db *sql.DB) repository.AssetRepository {
	return &AssetRepositoryPG{
		idGenerator,
		db,
	}
}

func (r *AssetRepositoryPG) AddAsset(payload *entity.AddAssetPayload) string {
	// Create ID
	id := r.idGenerator.Generate()

	// Query
	query := `INSERT INTO 
    			assets(id, title, file_path, file_watermark_path, description, details, price, owner_id) 
			  VALUES
			      ($1, $2, $3, $4, $5, $6, $7, $8)
			  RETURNING id`
	var returnedId string
	err := r.db.QueryRow(
		query,
		id,
		payload.Title,
		payload.OriginalLink,
		payload.WatermarkLink,
		payload.Description,
		payload.Details,
		payload.Price,
		payload.OwnerId,
	).Scan(&returnedId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("add_asset_err: %v", err))
	}

	return returnedId
}

func (r *AssetRepositoryPG) GetPreviewAssets(listCount int, pageList int, userId string) []entity.PreviewAsset {
	// Pagination
	offset := (pageList - 1) * listCount

	// Query
	query := `
			SELECT 
				a.id, 
				u.username AS owner_username, 
				a.title, 
				a.file_watermark_path, 
				a.description,
				COUNT(f.id) AS favorite_count,
				CASE WHEN uf.asset_id IS NOT NULL THEN true ELSE false END AS is_favorite
			FROM assets a
			INNER JOIN users u ON a.owner_id = u.id
			LEFT JOIN favorites f ON a.id = f.asset_id
			LEFT JOIN favorites uf ON a.id = uf.asset_id AND uf.user_id = $3
			GROUP BY a.id, u.username, a.title, a.file_watermark_path, a.description, a.created_at, uf.asset_id
			ORDER BY a.created_at DESC
			LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, listCount, offset, userId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("get_preview_assets_err: %v", err))
	}

	return services.GetTableDB[entity.PreviewAsset](rows)
}

func (r *AssetRepositoryPG) GetDetailAsset(id string, userId string) *entity.Asset {
	var result entity.Asset

	var originalPath string
	var watermarkPath string

	// Query
	query := `
			SELECT
				a.id,
				u.id AS owner_id,
				u.username AS owner_username, 
				a.title, 
				a.file_path,
				a.file_watermark_path,
				a.description,
				a.details, 
				a.price,
				a.created_at, 
				a.updated_at,
				COUNT(f.id) AS favorite_count,
				CASE WHEN uf.asset_id IS NOT NULL THEN true ELSE false END AS is_favorite
			FROM assets a
			INNER JOIN users u ON a.owner_id = u.id
			LEFT JOIN favorites f ON a.id = f.asset_id
			LEFT JOIN favorites uf ON a.id = uf.asset_id AND uf.user_id = $2
			WHERE a.id = $1
			GROUP BY a.id, u.id, u.username, a.title, a.file_path, a.file_watermark_path, a.description, a.details, a.price, a.created_at, a.updated_at, uf.asset_id`
	err := r.db.QueryRow(query, id, userId).Scan(
		&result.Id,
		&result.OwnerId,
		&result.OwnerUsername,
		&result.Title,
		&originalPath,
		&watermarkPath,
		&result.Description,
		&result.Details,
		&result.Price,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.FavoriteCount,
		&result.IsFavorite,
	)

	// Evaluate
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fiber.NewError(fiber.StatusNotFound, "Asset is not existed!"))
		}

		panic(fmt.Errorf("get_detail_asset_err: %v", err))
	}

	// If user is the owner, return the original instead
	if result.OwnerId == userId {
		result.FilePath = originalPath
	} else {
		result.FilePath = watermarkPath
	}
	fmt.Printf("user: %s; asset: %s", userId, result.FilePath)

	return &result
}

func (r *AssetRepositoryPG) UpdateAsset(id string, payload *entity.AddAssetPayload) (string, string) {
	var oldOriginalLink string
	var oldWatermarkLink string

	// Query
	query := `
			WITH old_data AS (
				SELECT file_path, file_watermark_path
				FROM assets
				WHERE id = $1
			)
			UPDATE assets
			SET 
				title = $2,
				file_path = $3,
			    file_watermark_path = $4,
				description = $5, 
				details = $6,
				price = $7,
				updated_at = NOW()
			FROM old_data
			WHERE id = $1
			RETURNING old_data.file_path, old_data.file_watermark_path`
	err := r.db.QueryRow(
		query,
		id,
		payload.Title,
		payload.OriginalLink,
		payload.WatermarkLink,
		payload.Description,
		payload.Details,
		payload.Price,
	).Scan(&oldOriginalLink, &oldWatermarkLink)

	// Evaluate
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fiber.NewError(fiber.StatusNotFound, "Asset is not existed!"))
		}

		panic(fmt.Errorf("update_asset_err: %v", err))
	}

	return oldOriginalLink, oldWatermarkLink
}

func (r *AssetRepositoryPG) DeleteAsset(id string) (string, string) {
	var oldOriginalLink string
	var oldWatermarkLink string

	// Query
	query := `DELETE FROM assets WHERE id = $1 RETURNING file_path, file_watermark_path`
	err := r.db.QueryRow(query, id).Scan(&oldOriginalLink, &oldWatermarkLink)

	// Evaluate
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fiber.NewError(fiber.StatusNotFound, "Asset is not existed!"))
		}

		panic(fmt.Errorf("delete_asset_err: %v", err))
	}

	return oldOriginalLink, oldWatermarkLink
}

func (r *AssetRepositoryPG) VerifyOwner(userId string, id string) {
	var ownerId string

	// Query
	query := "SELECT owner_id FROM assets WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&ownerId)

	// Evaluate
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fiber.NewError(fiber.StatusNotFound, "Asset is not existed!"))
		}

		panic(fmt.Errorf("verify_owner_err: %v", err))
	}

	if userId != ownerId {
		panic(fiber.NewError(fiber.StatusForbidden, "You don't have permission to do the action!"))
	}
}
