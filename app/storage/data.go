package storage

import (
	"github.com/satioO/go_practices/app/models"
)

// PopulateBeers populates the products
func PopulateBeers() {
	beers := []models.Beer{
		models.Beer{
			ID:    1,
			Name:  "Pliny the Elder",
			Price: 100,
		},
		models.Beer{
			ID:    2,
			Name:  "Oatmeal Stout",
			Price: 200,
		},
		models.Beer{
			ID:    3,
			Name:  "Märzen",
			Price: 300,
		},
	}

	DB.SaveBeers(beers)
}
