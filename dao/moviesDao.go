package dao

import (
	. "github.com/sil-vio/golang-restapi/models"
)

type MoviesDAOInterface interface {
	FindAll() ([]Movie, error)
	FindById(id string) (Movie, error)
	Insert(movie Movie) (string, error)
	Delete(id string) error
	Update(movie Movie) error
}
