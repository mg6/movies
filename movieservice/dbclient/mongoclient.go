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

func (m *MongoClient) CreateMovie(movie *model.Movie) error {
  err := m.Session.DB("app").C("movies").Insert(&movie)
  return err
}

func (m *MongoClient) GetMovies() (model.Movies, error) {
  var movies model.Movies
  err := m.Session.DB("app").C("movies").Find(nil).All(&movies)
  if err != nil {
    log.Println(err)
    return nil, err
  }
  return movies, nil
}

func (m *MongoClient) DeleteMovie(id string) error {
  err := m.Session.DB("app").C("movies").Remove(bson.M{"_id": bson.ObjectIdHex(id)})
  return err
}

func (m *MongoClient) CreateReview(movieId string, review *model.Review) error {
  var movie model.Movie
  err := m.Session.DB("app").C("movies").Find(bson.M{"_id": bson.ObjectIdHex(movieId)}).One(&movie)
  if err != nil {
    log.Println(err)
    return err
  }

  movie.Reviews = append(movie.Reviews, *review)

  _, err = m.Session.DB("app").C("movies").Upsert(bson.M{"_id": bson.ObjectIdHex(movieId)}, &movie)
  return err
}

func (m *MongoClient) GetReviews(movieId string) (model.Reviews, error) {
  var movie model.Movie
  err := m.Session.DB("app").C("movies").Find(bson.M{"_id": bson.ObjectIdHex(movieId)}).One(&movie)
  if err != nil {
    log.Println(err)
    return nil, err
  }
  return movie.Reviews, nil
}
