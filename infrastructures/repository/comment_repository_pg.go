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

type CommentRepositoryPG struct {
	db          *sql.DB
	idGenerator generator.IdGenerator
}

func NewCommentRepositoryPG(db *sql.DB, idGenerator generator.IdGenerator) repository.CommentRepository {
	return &CommentRepositoryPG{
		db:          db,
		idGenerator: idGenerator,
	}
}

func (r *CommentRepositoryPG) CreateComment(payload *entity.CreateCommentPayload) string {
	// Create ID
	id := r.idGenerator.Generate()

	// Query
	query := `
		INSERT INTO 
		    comments(id, asset_id, user_id, content) 
		VALUES 
		    ($1, $2, $3, $4)
		RETURNING id
	`

	var returnedId string
	err := r.db.QueryRow(
		query,
		id,
		payload.AssetId,
		payload.UserId,
		payload.Content,
	).Scan(&returnedId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("comment_repo_pg_error: create comment: %v", err))
	}

	return returnedId
}

func (r *CommentRepositoryPG) GetCommentsByAsset(assetId string) []entity.Comment {
	// Query
	query := `
		SELECT 
		    c.id, 
		    c.asset_id, 
		    c.user_id, 
		    u.username, 
		    u.avatar_link,
		    c.content, 
		    c.created_at, 
		    c.updated_at 
		FROM comments c
		INNER JOIN users u ON u.id = c.user_id
		WHERE c.asset_id = $1
	`

	rows, err := r.db.Query(query, assetId)
	if err != nil {
		panic(fmt.Errorf("comment_repo_pg_error: get comments by asset: %v", err))
	}

	return services.GetTableDB[entity.Comment](rows)
}

func (r *CommentRepositoryPG) UpdateComment(id string, content string) {
	// Query
	query := `
		UPDATE comments 
		SET content = $2, updated_at = now() 
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id, content)
	if err != nil {
		panic(fmt.Errorf("comment_repo_pg_error: update comment: %v", err))
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		panic(fiber.NewError(fiber.StatusNotFound, "Comment not found!"))
	}
}

func (r *CommentRepositoryPG) DeleteComment(id string) {
	// Query
	query := `DELETE FROM comments WHERE id = $1`
	result, err := r.db.Exec(query, id)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("comment_repo_pg_error: delete comment: %v", err))
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		panic(fiber.NewError(fiber.StatusNotFound, "Comment not found!"))
	}
}
