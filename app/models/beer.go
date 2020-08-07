package models

import (
	"errors"
)

// Beer represents the beer unit
type Beer struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Price   int16    `json:"price"`
	Reviews []Review `json:"reviews"`
}

// Beers represent collection of Beer
type Beers []Beer

// Find returns the index of review
func (b *Beers) Find(fn func(Beer, int) bool) (Beer, error) {
	for i, val := range *b {
		if cond := fn(val, i); cond == true {
			return val, nil
		}
	}

	return Beer{}, errors.New("Beer does not exists")
}
