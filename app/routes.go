package app

import "github.com/satioO/go_practices/app/handlers"

// SetupRoutes defines all routes
func (a *App) SetupRoutes() {
	// Beer Routes
	a.Router.HandleFunc("/beers", handlers.GetBeers).Methods("GET")
	a.Router.HandleFunc("/beers/{beerID}", handlers.GetBeer).Methods("GET")

	// Reviews Routes
	a.Router.HandleFunc("/beers/{beerID}/reviews", handlers.GetReviews).Methods("GET")
}
