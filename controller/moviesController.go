package controller

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/sil-vio/golang-restapi/daoImpl"
	. "github.com/sil-vio/golang-restapi/models"

	"github.com/gorilla/mux"

	. "github.com/sil-vio/golang-restapi/services"
	. "github.com/sil-vio/golang-restapi/servicesImpl"
)

//
type MoviesController struct {
	moviesService IMovieService
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func MovieControllerFactory() MoviesController {
	log.Println("MovieCotrollerFacory -- START --")
	var moviesController = MoviesController{MoviesServicesFactory(MoviesDAOFactory())}
	log.Printf("movieController: %#+v\n", moviesController)
	return moviesController
}

// GET list of movies
func (c MoviesController) AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	movies, err := c.moviesService.GetAllMovies()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

// GET a movie by its ID
func (c MoviesController) FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := c.moviesService.GetMovieById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

// POST a new movie
func (c MoviesController) CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	id, err := c.moviesService.InsertMovie(movie)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("film creato con id: %#+v\n", id)
	movie, _ = c.moviesService.GetMovieById(id)
	respondWithJson(w, http.StatusCreated, movie)
}

// PUT update an existing movie
func (c MoviesController) UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := c.moviesService.UpdateMovie(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing movie
func (c MoviesController) DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Printf("id: %#+v\n", id)
	if err := c.moviesService.DeleteMovieById(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
