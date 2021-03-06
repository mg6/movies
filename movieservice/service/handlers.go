package service

import (
  "encoding/json"
  "log"
  "net/http"
  "sort"
  "time"
  "github.com/gorilla/mux"
  "github.com/mg6/movies/movieservice/model"
  "github.com/mg6/movies/movieservice/dbclient"
  "github.com/mg6/movies/movieservice/approvalclient"
)

var DbClient dbclient.DbClient
var ApprovalClient approvalclient.ApprovalClient

func CreateMovie(w http.ResponseWriter, r *http.Request) {
  var m model.Movie
  if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  if ok := m.IsValid(); !ok {
    http.Error(w, "Validation unsuccessful", http.StatusBadRequest)
    return
  }

  m.CreatedAt = time.Now()
  created, err := DbClient.CreateMovie(&m)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  if err := json.NewEncoder(w).Encode(created); err != nil {
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
  vars := mux.Vars(r)
  id, ok := vars["movieId"]
  if !ok {
    http.Error(w, "Missing movie ID", http.StatusBadRequest)
    return
  }

  err := DbClient.DeleteMovie(id)
  if err != nil {
    if err, ok := err.(*dbclient.ErrNotFound); ok {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
    } else {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  w.WriteHeader(http.StatusNoContent)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  movieId, ok := vars["movieId"]
  if !ok {
    http.Error(w, "Missing movie ID", http.StatusBadRequest)
    return
  }

  var review model.Review
  if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  review.Status = model.Unapproved
  review.CreatedAt = time.Now()

  created, err := DbClient.CreateReview(movieId, &review)
  if err != nil {
    if err, ok := err.(*dbclient.ErrNotFound); ok {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
    } else {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  go func() {
    log.Println("Requesting review approval")
    reply, err := ApprovalClient.RequestApproval(*created)
    if err != nil {
      log.Printf("Cannot get approval reply: %s", err)
      return
    }

    time.Sleep(1 * time.Second)   // simulate delayed approval
    log.Printf("Got approval response with status %v\n", reply.Status)
    created.Status = reply.Status

    if reply.Status == model.Approved {
      DbClient.ApproveReview(movieId, created)
    }
  }()

  w.WriteHeader(http.StatusCreated)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  if err := json.NewEncoder(w).Encode(created); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func GetReviews(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  movieId, ok := vars["movieId"]
  if !ok {
    http.Error(w, "Missing movie ID", http.StatusBadRequest)
    return
  }

  reviews, err := DbClient.GetReviews(movieId)
  if err != nil {
    if err, ok := err.(*dbclient.ErrNotFound); ok {
      http.Error(w, err.Error(), http.StatusNotFound)
      return
    } else {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }

  result := make(model.Reviews, len(reviews))
  for i, movie := range reviews {
    result[i] = movie
  }

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(result); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
