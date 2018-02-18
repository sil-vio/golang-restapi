package main

import (
	"log"
	"net/http"

	. "github.com/sil-vio/golang-restapi/controller"

	"github.com/gorilla/mux"
)

var movieController = MovieControllerFactory()

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", movieController.AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies", movieController.CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", movieController.UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies/{id}", movieController.DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", movieController.FindMovieEndpoint).Methods("GET")
	log.Println("Avvio server in ascolto su porta 3000!")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
