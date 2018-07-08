package dbclient

import (
  "github.com/mg6/movies/movieservice/model"
)

type DbClient interface {
  CreateMovie(*model.Movie) error
  GetMovies() (model.Movies, error)
  DeleteMovie(string) error
}
