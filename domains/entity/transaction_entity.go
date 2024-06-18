package entity

import "time"

// CreateTransactionPayload is the payload for creating a transaction
type CreateTransactionPayload struct {
	UserID      string   `json:"user_id"`
	TotalAmount int      `json:"total_amount"`
	AssetsId    []string `json:"items_id"`
}

type PreviewTransactionItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
}

// TransactionItem represents an item within a transaction
type TransactionItem struct {
	ID           string    `json:"id"`
	AssetID      string    `json:"asset_id"`
	Title        string    `json:"title"`
	OriginalPath string    `json:"original_path"`
	Description  string    `json:"description"`
	PurchasedAt  time.Time `json:"purchased_at"`
}

type PreviewTransaction struct {
	Id          string `json:"id"`
	TotalAmount string `json:"total_amount"`
	PurchasedAt string `json:"purchased_at"`
}

// Transaction represents a transaction record
type Transaction struct {
	ID          string                   `json:"id"`
	TotalAmount int64                    `json:"total_amount"`
	PurchasedAt time.Time                `json:"purchased_at"`
	Items       []PreviewTransactionItem `json:"items"` // List of items in the transaction
}
