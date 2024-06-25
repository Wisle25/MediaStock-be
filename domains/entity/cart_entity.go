package entity

// CartPayload represents the data required to add or remove an item from the cart.
type CartPayload struct {
	AssetId string `json:"assetId"` // ID of the asset to be added or removed from the cart
	UserId  string `json:"userId"`  // ID of the user who owns the cart
}

// Cart represents the details of an item in the user's cart.
type Cart struct {
	Id       string `json:"id"`       // Unique identifier for the cart item
	Title    string `json:"title"`    // Title of the asset in the cart
	Price    string `json:"price"`    // Price of the asset in the cart
	FilePath string `json:"filePath"` // Path to the asset file
}
