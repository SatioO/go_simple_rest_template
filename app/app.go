package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/satioO/go_practices/app/handlers"
	"github.com/satioO/go_practices/app/storage"
)

// App defines the routing configuration
type App struct {
	Router *mux.Router
}

// Initialize the storage and router
func (a *App) Initialize() {
	storage.AttachStorage()
	storage.PopulateBeers()

	a.Router = mux.NewRouter()
}

// Run the application
func (a *App) Run(addr string) {
	fmt.Println("The beer server is on tap now.")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// SetupRoutes has the list of routes
func (a *App) SetupRoutes() {
	// ROUTES: BEERS
	a.Router.HandleFunc("/beers", handlers.GetBeers).Methods("GET")
	a.Router.HandleFunc("/beers/{beerId}", handlers.GetBeer).Methods("GET")

	// ROUTES: REVIEWS
	a.Router.HandleFunc("/beers/{beerId}/reviews", handlers.GetReviews).Methods("GET")
}
