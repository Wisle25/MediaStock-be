﻿package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/wisle25/media-stock-be/applications/generator"
	"github.com/wisle25/media-stock-be/commons"
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

func (r *AssetRepositoryPG) AddAsset(payload *entity.AssetPayload) string {
	// Create ID
	id := r.idGenerator.Generate()

	// Query
	query := `INSERT INTO 
    			assets(id, title, file_path, file_watermark_path, category, description, details, price, owner_id) 
			  VALUES
			      ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			  RETURNING id`
	var returnedId string
	err := r.db.QueryRow(
		query,
		id,
		payload.Title,
		payload.OriginalLink,
		payload.WatermarkLink,
		payload.Category,
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

func (r *AssetRepositoryPG) GetPreviewAssets(listCount int, pageList int, userId string, sortBy string, search string, category string) []entity.PreviewAsset {
	// Pagination
	offset := (pageList - 1) * listCount

	// Determine sorting column and order
	var sortColumn string
	switch sortBy {
	case "Recommended":
		sortColumn = "rating DESC, a.created_at DESC"
	case "Newest":
		sortColumn = "a.created_at DESC"
	case "Oldest":
		sortColumn = "a.created_at ASC"
	case "PriceLow":
		sortColumn = "a.price ASC"
	case "PriceHigh":
		sortColumn = "a.price DESC"
	default:
		sortColumn = "a.created_at DESC"
	}

	// Query
	query := fmt.Sprintf(`
		SELECT 
			a.id, 
			u.username AS owner_username, 
			a.title, 
			a.file_watermark_path,
			a.category,
			a.description,
			COALESCE(AVG(r.score), 0) AS rating,
			COUNT(f.id) AS favorite_count,
			CASE WHEN uf.asset_id IS NOT NULL THEN true ELSE false END AS is_favorite
		FROM assets a
		INNER JOIN users u ON a.owner_id = u.id
		LEFT JOIN favorites f ON a.id = f.asset_id
		LEFT JOIN favorites uf ON a.id = uf.asset_id AND uf.user_id = $3
		LEFT JOIN ratings r ON a.id = r.asset_id
		WHERE ($4 = '' OR a.title ILIKE '%%' || $4 || '%%' OR a.description ILIKE '%%' || $4 || '%%')
		AND ($5 = '' OR a.category = $5)
		GROUP BY a.id, u.username, a.title, a.file_watermark_path, a.description, a.created_at, uf.asset_id
		ORDER BY %s
		LIMIT $1 OFFSET $2`, sortColumn)

	rows, err := r.db.Query(query, listCount, offset, userId, search, category)
	// Evaluate
	if err != nil {
		commons.ThrowServerError("get_preview_assets_err:", err)
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
				a.category,
				a.description,
				a.details, 
				a.price,
				a.created_at, 
				a.updated_at,
				COUNT(DISTINCT f.id) AS favorite_count,
				COUNT(DISTINCT ti.id) AS purchased_count,
				CASE WHEN uf.asset_id IS NOT NULL THEN true ELSE false END AS is_favorite,
				CASE WHEN EXISTS (
					SELECT 1 
					FROM transaction_items ti_user
					JOIN transactions t_user ON ti_user.transaction_id = t_user.id
					WHERE ti_user.asset_id = a.id AND t_user.user_id = $2
				) THEN true ELSE false END AS is_purchased,
				COALESCE(AVG(r.score), 0) AS rating
			FROM assets a
			INNER JOIN users u ON a.owner_id = u.id
			LEFT JOIN favorites f ON a.id = f.asset_id
			LEFT JOIN favorites uf ON a.id = uf.asset_id AND uf.user_id = $2
			LEFT JOIN transaction_items ti ON a.id = ti.asset_id
			LEFT JOIN transactions t ON t.id = ti.transaction_id
			LEFT JOIN ratings r ON a.id = r.asset_id
			WHERE a.id = $1
			GROUP BY a.id, u.id, u.username, a.title, a.file_path, a.file_watermark_path, a.description, a.details, a.price, a.created_at, a.updated_at, uf.asset_id`
	err := r.db.QueryRow(query, id, userId).Scan(
		&result.Id,
		&result.OwnerId,
		&result.OwnerUsername,
		&result.Title,
		&originalPath,
		&watermarkPath,
		&result.Category,
		&result.Description,
		&result.Details,
		&result.Price,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.FavoriteCount,
		&result.PurchasedCount,
		&result.IsFavorite,
		&result.IsPurchased,
		&result.Rating,
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

	return &result
}

func (r *AssetRepositoryPG) GetPurchasedAsset(userId string) []entity.SimpleAsset {
	// Query
	query := `
			SELECT
    			a.id,
				a.title,
				a.price,
			    a.file_watermark_path AS file_path
			FROM
				assets a
			INNER JOIN transaction_items ON a.id = transaction_items.asset_id
			INNER JOIN transactions ON transaction_items.transaction_id = transactions.id
			WHERE transactions.user_id = $1`
	rows, err := r.db.Query(query, userId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("get_purchased_asset_err: %v", err))
	}

	return services.GetTableDB[entity.SimpleAsset](rows)
}

func (r *AssetRepositoryPG) GetOwnedAsset(userId string) []entity.SimpleAsset {
	// Query
	query := `
			SELECT
    			id,
				title,
				price,
			    file_watermark_path AS file_path
			FROM
				assets
			WHERE assets.owner_id = $1`
	rows, err := r.db.Query(query, userId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("get_owned_asset_err: %v", err))
	}

	return services.GetTableDB[entity.SimpleAsset](rows)
}

func (r *AssetRepositoryPG) UpdateAsset(id string, payload *entity.AssetPayload) (string, string) {
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

func (r *AssetRepositoryPG) DownloadAsset(id string, userId string) (string, string) {
	var title string
	var originalLink string
	var isPurchased bool

	// Query
	query := `
		SELECT 
		    a.title,
			a.file_path,
			CASE WHEN ti.asset_id IS NOT NULL THEN true ELSE false END AS is_purchased
		FROM assets a
		LEFT JOIN transaction_items ti ON a.id = ti.asset_id
		LEFT JOIN transactions t ON t.id = ti.transaction_id AND t.user_id = $2
		WHERE a.id = $1`
	err := r.db.QueryRow(query, id, userId).Scan(&title, &originalLink, &isPurchased)

	// Evaluate
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			panic(fiber.NewError(fiber.StatusNotFound, "Asset is not existed!"))
		}
		panic(fmt.Errorf("download_asset_err: %v", err))
	}

	if !isPurchased {
		panic(fiber.NewError(fiber.StatusForbidden, "You haven't purchased the asset!"))
	}

	return title, originalLink
}
