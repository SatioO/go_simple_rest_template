package storage

import (
	"github.com/satioO/go_practices/app/models"
)

// Storage defines the functionality of a data store for the beer service
type Storage interface {
	FindBeers() models.Beers
	SaveBeers(models.Beers)
	FindBeer(beerID int) models.Beer
	FindReviews(beerID int) models.Reviews
}

// DB is the "global" storage instance
var DB Storage

// AttachStorage creates new storage
func AttachStorage() {
	DB = new(MemoryStorage)
}
