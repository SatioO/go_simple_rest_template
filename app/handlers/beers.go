package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/satioO/go_practices/app/storage"
	"github.com/satioO/go_practices/app/utils"
)

// GetBeers gets all the beers in the inventory
func GetBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cellar := storage.DB.FindBeers()
	utils.RespondWithJSON(w, http.StatusOK, cellar)
}

// GetBeer gets all the beer in the inventory
func GetBeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	beerID, _ := strconv.ParseInt(vars["beerID"], 0, 64)
	cellar := storage.DB.FindBeer(int(beerID))
	utils.RespondWithJSON(w, http.StatusOK, cellar)
}
