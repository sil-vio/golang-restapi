package servicesImpl

import (
	"errors"

	. "github.com/sil-vio/golang-restapi/dao"
	. "github.com/sil-vio/golang-restapi/models"
	. "github.com/sil-vio/golang-restapi/services"

	"gopkg.in/mgo.v2/bson"
)

func MoviesServicesFactory(dao MoviesDAOInterface) IMovieService {
	return moviesServicesImpl{movieDao: dao}
}

type moviesServicesImpl struct {
	movieDao MoviesDAOInterface
}

func (ms moviesServicesImpl) GetAllMovies() ([]Movie, error) {
	movies, err := ms.movieDao.FindAll()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (ms moviesServicesImpl) GetMovieById(id string) (Movie, error) {
	movie, err := ms.movieDao.FindById(id)
	if err != nil {
		return Movie{}, err
	}
	return movie, nil
}

func (ms moviesServicesImpl) InsertMovie(movie Movie) (string, error) {
	if movie.Name == "" {
		return "", errors.New("Nome film non valorizzato!")
	}
	id, err := ms.movieDao.Insert(movie)
	if err != nil {
		return "", err
	}
	return id, nil

}

func (ms moviesServicesImpl) UpdateMovie(movie Movie) error {

	movie.ID = bson.NewObjectId()
	if movie.ID == "" || movie.Name == "" {
		return errors.New("Film non valido!")
	}
	return ms.movieDao.Update(movie)

}

func (ms moviesServicesImpl) DeleteMovieById(id string) error {

	return ms.movieDao.Delete(id)

}
