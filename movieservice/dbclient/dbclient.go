package dbclient

import (
  "github.com/mg6/movies/movieservice/model"
)

type DbClient interface {
  Connect(url string) error
  CreateMovie(*model.Movie) (*model.Movie, error)
  GetMovies() (model.Movies, error)
  DeleteMovie(string) error
  CreateReview(movieId string, review *model.Review) (*model.Review, error)
  ApproveReview(movieId string, review *model.Review) error
  GetReviews(movieId string) (model.Reviews, error)
}

type ErrNotFound struct {}

func (e *ErrNotFound) Error() string {
  return "Entity not found"
}
