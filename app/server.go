package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/satioO/go_practices/app/storage"
)

// Server defines the routing configuration
type Server struct {
	DB     *storage.Storage
	Router *mux.Router
}

// Initialize the storage and router
func (a *Server) Initialize() {
	storage.AttachStorage()
	storage.PopulateBeers()

	a.Router = mux.NewRouter()
	a.DB = &storage.DB

	a.SetupRoutes()
}

// Run the application
func (a *Server) Run(addr string) {
	fmt.Println("The beer server is on tap now.")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
