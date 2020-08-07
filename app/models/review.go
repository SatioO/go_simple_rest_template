package models

import "errors"

// Review represents the review on beer
type Review struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

// Reviews represent collection of Beer
type Reviews []Review

// Find returns the index of review
func (r *Reviews) Find(fn func(Review, int) bool) (Review, error) {
	for i, val := range *r {
		if cond := fn(val, i); cond == true {
			return val, nil
		}
	}

	return Review{}, errors.New("Beer does not exists")
}
