package storage

import (
	"github.com/satioO/go_practices/app/models"
)

// MemoryStorage layered save only in memory
type MemoryStorage struct {
	cellar []models.Beer
}

// FindBeers return all beers
func (s *MemoryStorage) FindBeers() []models.Beer {
	return s.cellar
}

// SaveBeers return all beers
func (s *MemoryStorage) SaveBeers(beers []models.Beer) {
	s.cellar = beers
}

// FindBeer returns a beer
func (s *MemoryStorage) FindBeer(beerID int) models.Beer {
	beer := models.Beer{}

	return beer
}

// FindReviews returns a reviews on beer
func (s *MemoryStorage) FindReviews(beerID int) []models.Review {
	return []models.Review{}
}
