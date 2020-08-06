package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/satioO/go_practices/app/storage"
)

// GetBeers gets all the beers in the inventory
func GetBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cellar := storage.DB.FindBeers()
	json.NewEncoder(w).Encode(cellar)
}

// GetBeer gets all the beer in the inventory
func GetBeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	w.Header().Set("Content-Type", "application/json")
	cellar := storage.DB.FindBeer(int(id))
	json.NewEncoder(w).Encode(cellar)
}
