package model

import (
  "time"
  "github.com/globalsign/mgo/bson"
)

type Movie struct {
  Id          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
  Title       string          `json:"title"`
  Rating      float64         `json:"rating"`
  Director    string          `json:"director"`
  Actors      []string        `json:"actors"`
  CreatedAt   time.Time       `json:"createdAt"`
}

type Movies []Movie
