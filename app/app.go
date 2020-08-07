package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

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
	a.SetupRoutes()
}

// Run the application
func (a *App) Run(addr string) {
	fmt.Println("The beer server is on tap now.")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
