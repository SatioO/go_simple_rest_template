package handlers

import (
	"net/http"
	"strconv"

	"github.com/satioO/go_practices/app/utilities"

	"github.com/gorilla/mux"
	"github.com/satioO/go_practices/app/storage"
)

// GetReviews gets all reviews of a beer
func GetReviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	beerID, _ := strconv.ParseInt(vars["beerID"], 0, 64)

	reviews := storage.DB.FindReviews(int(beerID))

	utilities.RespondWithJSON(w, http.StatusOK, reviews)
}
