package storage

import (
	"github.com/satioO/go_practices/app/models"
)

// Storage defines the functionality of a data store for the beer service
type Storage interface {
	FindBeers() []models.Beer
	SaveBeers([]models.Beer)
	FindBeer(id int) models.Beer
}

// DB is the "global" storage instance
var DB Storage

// AttachStorage creates new storage
func AttachStorage() {
	DB = new(MemoryStorage)
}
