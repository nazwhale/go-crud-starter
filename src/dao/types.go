package dao

type Work struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	ArtistID int    `json:"artist_id"`
}
