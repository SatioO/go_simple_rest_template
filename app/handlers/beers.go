package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/satioO/go_practices/app/storage"
	"github.com/satioO/go_practices/app/utilities"
)

// GetBeers gets all the beers in the inventory
func GetBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cellar := storage.DB.FindBeers()
	utilities.RespondWithJSON(w, http.StatusOK, cellar)
}

// GetBeer gets all the beer in the inventory
func GetBeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 0, 64)
	cellar := storage.DB.FindBeer(int(id))
	utilities.RespondWithJSON(w, http.StatusOK, cellar)
}
