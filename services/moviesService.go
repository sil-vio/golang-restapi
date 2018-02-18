package services

import (
	. "github.com/sil-vio/golang-restapi/models"
)

type IMovieService interface {
	GetAllMovies() ([]Movie, error)
	GetMovieById(id string) (Movie, error)
	InsertMovie(movie Movie) (string, error)
	DeleteMovieById(id string) error
	UpdateMovie(movie Movie) error
}
