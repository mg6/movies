package service


import (
  "net/http"
)

type Route struct {
  Name string
  Method string
  Pattern string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "CreateMovie",
    "POST",
    "/movies",
    CreateMovie,
  },
  Route{
    "GetAccount",
    "GET",
    "/movies",
    GetMoviesByRating,
  },
  Route{
    "DeleteMovie",
    "DELETE",
    "/movies/{movieId}",
    DeleteMovie,
  },
  Route{
    "CreateReview",
    "POST",
    "/reviews/{movieId}",
    CreateReview,
  },
}
