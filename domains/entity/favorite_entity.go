package entity

// FavoritePayload represents the data required to add or remove a favorite asset.
type FavoritePayload struct {
	AssetId string `json:"assetId"` // ID of the asset to be favorited
	UserId  string `json:"userId"`  // ID of the user who is favoriting the asset
}

// Favorite represents the details of a favorited asset.
type Favorite struct {
	Id       string `json:"id"`       // Unique identifier for the favorite item
	Title    string `json:"title"`    // Title of the favorited asset
	Price    string `json:"price"`    // Price of the favorited asset
	FilePath string `json:"filePath"` // Path to the favorited asset file
}
