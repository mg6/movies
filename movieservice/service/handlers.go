package service

import (
  "encoding/json"
  "net/http"
  "sort"
  "time"
  "github.com/mg6/movies/movieservice/model"
  "github.com/mg6/movies/movieservice/dbclient"
)

var DbClient dbclient.DbClient

func CreateMovie(w http.ResponseWriter, r *http.Request) {
  var m model.Movie
  if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  m.CreatedAt = time.Now()
  if err := DbClient.CreateMovie(&m); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  if err := json.NewEncoder(w).Encode(&m); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func GetMoviesByRating(w http.ResponseWriter, r *http.Request) {
  movies, err := DbClient.GetMovies()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  result := make(model.Movies, len(movies))
  for i, movie := range movies {
    result[i] = movie
  }

  sort.Sort(sort.Reverse(model.ByRating(result)))

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(result); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  json.NewEncoder(w).Encode([]interface{}{})
}
