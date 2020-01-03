package dao

type Artist struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

type Work struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	ArtistID int    `json:"artist_id"`
}
