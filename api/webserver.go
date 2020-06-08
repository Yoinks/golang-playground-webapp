package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

//InitRouter Initializes the listening for url paths
func initRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/pokemans", GetPokemans).Methods("GET")

	return r

}

//InitWebServer Initializes the webserver listener and routes
func InitWebServer() {
	http.ListenAndServe(":8080", initRouter())
}
