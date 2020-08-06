package models

// Beer represents the beer unit
type Beer struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Price   int16    `json:"price"`
	Reviews []Review `json:"reviews"`
}
