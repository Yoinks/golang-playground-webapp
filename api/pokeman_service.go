package api

import (
	"encoding/json"
	"net/http"
)

//GetPokemans fetches Pokemans from database
func GetPokemans(w http.ResponseWriter, r *http.Request) {
	//TODO implement fetch from database
	json.NewEncoder(w).Encode("Testing")
}
