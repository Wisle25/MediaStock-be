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

type CartRepositoryPG struct {
	idGenerator generator.IdGenerator
	db          *sql.DB
}

func NewCartRepositoryPG(idGenerator generator.IdGenerator, db *sql.DB) repository.CartRepository {
	return &CartRepositoryPG{
		idGenerator,
		db,
	}
}

func (t *CartRepositoryPG) AddToCart(payload *entity.CartPayload) {
	id := t.idGenerator.Generate()

	// Query
	query := "INSERT INTO carts(id, asset_id, user_id) VALUES ($1, $2, $3)"
	result, err := t.db.Exec(query, id, payload.AssetId, payload.UserId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("add_favorite_err: %v", err))
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		panic(fiber.NewError(fiber.StatusBadRequest, "Unable to add cart"))
	}
}

func (t *CartRepositoryPG) GetAllCarts(userId string) []entity.Cart {
	// Query
	query := `
			SELECT 
    			a.id,
				a.title,
				a.price,
			    a.file_watermark_path AS file_path
			FROM carts f
			INNER JOIN assets a ON f.asset_id = a.id
			WHERE user_id = $1`
	rows, err := t.db.Query(query, userId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("get_all_carts_err: %v", err))
	}

	return services.GetTableDB[entity.Cart](rows)
}

func (t *CartRepositoryPG) RemoveCart(payload *entity.CartPayload) {
	// Query
	query := "DELETE FROM carts WHERE asset_id = $1 AND user_id = $2"
	result, err := t.db.Exec(query, payload.AssetId, payload.UserId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("remove_cart_err: %v", err))
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		panic(fiber.NewError(fiber.StatusBadRequest, "Unable to remove cart"))
	}
}

func (t *CartRepositoryPG) RemoveAllCartByUser(userId string) {
	// Query
	query := "DELETE FROM carts WHERE user_id = $1"
	result, err := t.db.Exec(query, userId)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("remove_cart_by_user_err: %v", err))
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		panic(fiber.NewError(fiber.StatusBadRequest, "Unable to remove all carts"))
	}
}
