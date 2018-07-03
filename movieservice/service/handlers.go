package service

import (
  "encoding/json"
  "net/http"
  "github.com/mg6/movies/movieservice/dbclient"
)

var DbClient dbclient.DbClient

func CreateMovie(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  json.NewEncoder(w).Encode([]interface{}{})
}

func GetMoviesByRating(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  json.NewEncoder(w).Encode([]interface{}{})
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  json.NewEncoder(w).Encode([]interface{}{})
}