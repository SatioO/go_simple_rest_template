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
func (s *MemoryStorage) FindBeer(id int) models.Beer {
	beer := models.Beer{}
	for i, num := range s.cellar {
		if num.ID == id {
			beer = s.cellar[i]
		}
	}
	return beer
}
