package model

import (
  "regexp"
  "time"
  "github.com/globalsign/mgo/bson"
)

type Movie struct {
  Id          bson.ObjectId   `json:"-" bson:"_id,omitempty"`
  Slug        string          `json:"id"`
  Title       string          `json:"title"`
  Rating      float64         `json:"rating"`
  Director    string          `json:"director"`
  Actors      []string        `json:"actors"`
  CreatedAt   time.Time       `json:"createdAt"`
}

type Movies []Movie


func (m *Movie) IsValid() bool {
  validTitle := regexp.MustCompile("^[A-Za-z]{3,50}$")
  return validTitle.MatchString(m.Title)
}


type ByRating Movies

func (s ByRating) Len() int {
  return len(s)
}

func (s ByRating) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}

func (s ByRating) Less(i, j int) bool {
  return s[i].Rating < s[j].Rating
}
