package entity

type FavoritePayload struct {
	AssetId string `json:"asset_id"`
	UserId  string `json:"user_id"`
}

type Favorite struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Price    string `json:"price"`
	FilePath string `json:"file_path"`
}
