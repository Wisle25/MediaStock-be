package repository

import (
	"database/sql"
	"fmt"
	"github.com/wisle25/media-stock-be/applications/generator"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
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
    			assets(id, title, file_path, file_watermark_path, description, details, owner_id) 
			  VALUES
			      ($1, $2, $3, $4, $5, $6, $7)
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
		payload.OwnerId,
	).Scan(&returnedId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("add_asset_err: %v", err))
	}
	
	return returnedId
}
