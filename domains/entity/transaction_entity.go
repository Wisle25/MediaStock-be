package entity

import "time"

// CreateTransactionPayload is the payload for creating a transaction.
type CreateTransactionPayload struct {
	UserID      string   `json:"userId"`      // ID of the user making the transaction
	TotalAmount int      `json:"totalAmount"` // Total amount of the transaction
	AssetsId    []string `json:"assetsId"`    // List of asset IDs being purchased
}

// PreviewTransactionItem represents a brief overview of an item within a transaction.
type PreviewTransactionItem struct {
	ID    string `json:"id"`    // Unique identifier for the item
	Title string `json:"title"` // Title of the item
	Price string `json:"price"` // Price of the item
}

// TransactionItem represents an item within a transaction.
type TransactionItem struct {
	ID           string    `json:"id"`           // Unique identifier for the transaction item
	AssetID      string    `json:"assetId"`      // ID of the asset being purchased
	Title        string    `json:"title"`        // Title of the asset
	OriginalPath string    `json:"originalPath"` // Path to the original asset file
	Description  string    `json:"description"`  // Description of the asset
	PurchasedAt  time.Time `json:"purchasedAt"`  // Timestamp when the asset was purchased
}

// PreviewTransaction represents a brief overview of a transaction.
type PreviewTransaction struct {
	Id          string `json:"id"`          // Unique identifier for the transaction
	TotalAmount string `json:"totalAmount"` // Total amount of the transaction
	PurchasedAt string `json:"purchasedAt"` // Timestamp when the transaction was made
}

// Transaction represents a transaction record.
type Transaction struct {
	ID          string                   `json:"id"`          // Unique identifier for the transaction
	TotalAmount int64                    `json:"totalAmount"` // Total amount of the transaction
	PurchasedAt time.Time                `json:"purchasedAt"` // Timestamp when the transaction was made
	Items       []PreviewTransactionItem `json:"items"`       // List of items in the transaction
}
