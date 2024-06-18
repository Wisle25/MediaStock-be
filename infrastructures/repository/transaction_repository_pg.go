package repository

import (
	"database/sql"
	"fmt"
	"github.com/wisle25/media-stock-be/applications/generator"
	"github.com/wisle25/media-stock-be/domains/entity"
	"github.com/wisle25/media-stock-be/domains/repository"
	"github.com/wisle25/media-stock-be/infrastructures/services"
	"strings"
)

type TransactionRepositoryPG struct {
	idGenerator generator.IdGenerator
	db          *sql.DB
}

func NewTransactionRepositoryPG(idGenerator generator.IdGenerator, db *sql.DB) repository.TransactionRepository {
	return &TransactionRepositoryPG{
		idGenerator,
		db,
	}
}

func (t *TransactionRepositoryPG) CreateTransaction(payload *entity.CreateTransactionPayload) string {
	tx, err := t.db.Begin()
	if err != nil {
		panic(fmt.Errorf("create_transaction_begin_err: %v", err))
	}

	// Transaction
	transactionID := t.idGenerator.Generate()
	var returnedId string

	transactionQuery := `
		INSERT INTO 
		    transactions(id, user_id, total_amount)
		VALUES 
		    ($1, $2, $3) 
		RETURNING id`
	err = tx.QueryRow(
		transactionQuery,
		transactionID,
		payload.UserID,
		payload.TotalAmount,
	).Scan(&returnedId)

	if err != nil {
		_ = tx.Rollback()
		panic(fmt.Errorf("create_transaction_insert_err: %v", err))
	}

	// Prepare batch insert for transaction items
	var values []string
	args := []interface{}{transactionID}
	for i, assetID := range payload.AssetsId {
		values = append(values, fmt.Sprintf("($1, $%d)", i+2))
		args = append(args, assetID)
	}
	query := fmt.Sprintf(`
		INSERT INTO 
		    transaction_items(transaction_id, asset_id)
		VALUES %s`, strings.Join(values, ", "))

	_, err = tx.Exec(query, args...)
	if err != nil {
		_ = tx.Rollback()
		panic(fmt.Errorf("create_transaction_items_batch_err: %v", err))
	}

	err = tx.Commit()
	if err != nil {
		panic(fmt.Errorf("create_transaction_commit_err: %v", err))
	}

	return returnedId
}
func (t *TransactionRepositoryPG) GetTransactionByID(transactionID string) *entity.Transaction {
	var transaction entity.Transaction

	// Query
	transactionQuery := `
		SELECT 
		    id,
		    total_amount,
		    purchased_at
		FROM transactions
		WHERE id = $1`
	err := t.db.QueryRow(transactionQuery, transactionID).Scan(
		&transaction.ID,
		&transaction.TotalAmount,
		&transaction.PurchasedAt,
	)
	if err != nil {
		panic(fmt.Errorf("get_transaction_by_id_err: %v", err))
	}

	// Transaction item
	itemsQuery := `
		SELECT 
			a.id, a.title, a.description 
		FROM transaction_items ti
		INNER JOIN assets a ON a.id = ti.asset_id 
		WHERE ti.transaction_id = $1`
	rows, err := t.db.Query(itemsQuery, transactionID)

	if err != nil {
		panic(fmt.Errorf("get_transaction_items_by_id_err: %v", err))
	}

	transaction.Items = services.GetTableDB[entity.PreviewTransactionItem](rows)

	return &transaction
}

func (t *TransactionRepositoryPG) GetTransactionsByUser(userID string) []entity.PreviewTransaction {
	// Query
	query := `
		SELECT 
			id,
			total_amount,
			purchased_at
		FROM transactions
		WHERE user_id = $1`
	rows, err := t.db.Query(query, userID)

	// Evaluate
	if err != nil {
		panic(fmt.Errorf("get_transaction_items_by_user: %v", err))
	}

	return services.GetTableDB[entity.PreviewTransaction](rows)
}
