package storage

import (
	"github.com/satioO/go_practices/app/models"
)

// PopulateBeers populates the products
func PopulateBeers() {
	beers := []models.Beer{
		models.Beer{
			ID:      1,
			Name:    "Pliny the Elder",
			Price:   100,
			Reviews: []models.Review{},
		},
		models.Beer{
			ID:    2,
			Name:  "Oatmeal Stout",
			Price: 200,
			Reviews: []models.Review{
				models.Review{
					ID:   1,
					Body: "Excellent",
				},
			},
		},
		models.Beer{
			ID:      3,
			Name:    "Märzen",
			Price:   300,
			Reviews: []models.Review{},
		},
	}

	DB.SaveBeers(beers)
}
