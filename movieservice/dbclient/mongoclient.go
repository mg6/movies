package dbclient

import (
  "github.com/globalsign/mgo"
  "github.com/globalsign/mgo/bson"
  "github.com/mg6/movies/movieservice/model"
  "log"
)

type MongoClient struct {
  Session *mgo.Session
}

func (m *MongoClient) Connect(url string) error {
  log.Printf("Connecting to MongoDB at %v ...", url)
  session, err := mgo.Dial(url)
  if err != nil {
	log.Fatal(err)
  }

  m.Session = session
  return err
}

func (m *MongoClient) CreateMovie(movie *model.Movie) (*model.Movie, error) {
  info, err := m.Session.DB("app").C("movies").UpsertId(nil, &movie)
  if info.UpsertedId != nil {
    movie.Id = info.UpsertedId.(bson.ObjectId)
  }
  return movie, err
}

func (m *MongoClient) GetMovies() (model.Movies, error) {
  var movies model.Movies
  err := m.Session.DB("app").C("movies").Find(nil).All(&movies)
  if err != nil {
    log.Printf("Cannot get movies: %v", err)
    return nil, err
  }
  return movies, nil
}

func (m *MongoClient) DeleteMovie(slug string) error {
  err := m.Session.DB("app").C("reviews").Remove(bson.M{"movie": slug})
  err = m.Session.DB("app").C("movies").Remove(bson.M{"slug": slug})
  return err
}

func (m *MongoClient) CreateReview(movieSlug string, review *model.Review) (*model.Review, error) {
  var movie model.Movie
  err := m.Session.DB("app").C("movies").Find(bson.M{"slug": movieSlug}).One(&movie)
  if err != nil {
    log.Printf("Create review: error for movie %v: %v", movieSlug, err)
    return nil, err
  }

  review.Movie = movieSlug

  info, err := m.Session.DB("app").C("reviews").UpsertId(nil, &review)
  if info.UpsertedId != nil {
    review.Id = info.UpsertedId.(bson.ObjectId)
  }
  return review, err
}

func (m *MongoClient) ApproveReview(movieSlug string, review *model.Review) error {
  log.Println("Review approving not implemented")
  return nil
}

func (m *MongoClient) GetReviews(movieSlug string) (model.Reviews, error) {
  var reviews model.Reviews
  err := m.Session.DB("app").C("reviews").Find(bson.M{"movie": movieSlug}).All(&reviews)
  if err != nil {
    log.Printf("Get reviews: error for movie %v: %v", movieSlug, err)
    return nil, err
  }
  return reviews, nil
}
