package storage

import (
	"github.com/satioO/go_practices/app/models"
)

// MemoryStorage layered save only in memory
type MemoryStorage struct {
	cellar models.Beers
}

// FindBeers return all beers
func (s *MemoryStorage) FindBeers() models.Beers {
	return s.cellar
}

// SaveBeers return all beers
func (s *MemoryStorage) SaveBeers(beers models.Beers) {
	s.cellar = beers
}

// FindBeer returns a beer
func (s *MemoryStorage) FindBeer(beerID int) models.Beer {
	beer, err := s.cellar.Find(func(beer models.Beer, i int) bool {
		return beer.ID == beerID
	})

	if err != nil {
		return models.Beer{}
	}

	return beer
}

// FindReviews returns a reviews on beer
func (s *MemoryStorage) FindReviews(beerID int) models.Reviews {
	beer, err := s.cellar.Find(func(beer models.Beer, i int) bool {
		return beer.ID == beerID
	})

	if err != nil {
		return []models.Review{}
	}

	return beer.Reviews
}
