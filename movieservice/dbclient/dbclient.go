package dbclient

import (
  "github.com/mg6/movies/movieservice/model"
)

type DbClient interface {
  CreateMovie(*model.Movie) error
  GetMovies() (model.Movies, error)
  DeleteMovie(string) error
}

type ErrNotFound struct {}

func (e *ErrNotFound) Error() string {
  return "Entity not found"
}
