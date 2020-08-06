package models

// Review represents the review on beer
type Review struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}
