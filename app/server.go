package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/satioO/go_practices/app/storage"
	"github.com/satioO/go_practices/config"
)

// Server defines the routing configuration
type Server struct {
	DB     *storage.Storage
	Router *mux.Router
	Config *config.Configurations
}

// Initialize the storage and router
func (a *Server) Initialize() {
	storage.AttachStorage()
	storage.PopulateBeers()

	a.Router = mux.NewRouter()
	a.Config = a.SetupConfig()
	a.DB = &storage.DB

	a.SetupRoutes()
}

// SetupConfig defines the default configuration
func (a *Server) SetupConfig() *config.Configurations {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &configuration
}

// Run the application
func (a *Server) Run(addr string) {
	fmt.Println("The beer server is on tap now.")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
