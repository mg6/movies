package service

import (
  "bytes"
  "encoding/json"
  "net/http/httptest"
  "testing"
  . "github.com/stretchr/testify/mock"
  . "github.com/smartystreets/goconvey/convey"
  "github.com/mg6/movies/movieservice/model"
  "github.com/mg6/movies/movieservice/dbclient"
)

func TestCreateMovie(t *testing.T) {
  mockRepo := &dbclient.MockClient{}
  mockRepo.On("CreateMovie", Anything).Return(nil)
  DbClient = mockRepo

  expected := model.Movie{
    Title: "Test 123",
    Rating: 0.0,
    Director: "Director 123",
    Actors: []string{"A", "B", "C"},
  }
  asJson, _ := json.Marshal(expected)

  Convey("Given a HTTP request to create a movie", t, func() {
    req := httptest.NewRequest("POST", "/movies", bytes.NewReader(asJson))
    resp := httptest.NewRecorder()

    Convey("When the request is handled by the router", func() {
      NewRouter().ServeHTTP(resp, req)

      Convey("Then the response should be 201", func() {
        So(resp.Code, ShouldEqual, 201)

        Convey("And entity values should match", func() {
          actual := model.Movie{}
          json.Unmarshal(resp.Body.Bytes(), &actual)

          So(actual.Title, ShouldEqual, expected.Title)
          So(actual.Rating, ShouldEqual, expected.Rating)
          So(actual.Director, ShouldEqual, expected.Director)
          So(actual.Actors, ShouldResemble, expected.Actors)
        })
      })
    })
  })
}

func TestGetMoviesByRating(t *testing.T) {
  m1 := model.Movie{
    Title: "Test 123",
    Rating: 0.0,
    Director: "Director 123",
    Actors: []string{"A", "B", "C"},
  }
  m2 := model.Movie{
    Title: "Test 456",
    Rating: 2.0,
    Director: "Director 456",
    Actors: []string{"A", "B", "C"},
  }

  expected := model.Movies{m2, m1}

  mockRepo := &dbclient.MockClient{}
  mockRepo.On("GetMovies").Return(model.Movies{m1, m2}, nil)
  DbClient = mockRepo

  Convey("Given a HTTP request to get all movies", t, func() {
    req := httptest.NewRequest("GET", "/movies", nil)
    resp := httptest.NewRecorder()

    Convey("When the request is handled by the router", func() {
      NewRouter().ServeHTTP(resp, req)

      Convey("Then the response should be 200", func() {
        So(resp.Code, ShouldEqual, 200)

        Convey("And entity list should be sorted descending by rating", func() {
          actual := model.Movies{}
          json.Unmarshal(resp.Body.Bytes(), &actual)

          So(actual, ShouldResemble, expected)
        })
      })
    })
  })
}

func TestGetWrongPath(t *testing.T) {
  Convey("Given a HTTP request for /invalid/123", t, func() {
    req := httptest.NewRequest("GET", "/invalid/123", nil)
    resp := httptest.NewRecorder()

    Convey("When the request is handled by the router", func() {
      NewRouter().ServeHTTP(resp, req)

      Convey("Then the response should be a 404", func() {
        So(resp.Code, ShouldEqual, 404)
      })
    })
  })
}
