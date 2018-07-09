package model

import (
  "time"
  "github.com/globalsign/mgo/bson"
)

type Status string

const (
  Unapproved  Status = "unapproved"
  Approved    Status = "approved"
)

type Review struct {
  Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Movie     string        `json:"movie"`
  Text      string        `json:"text"`
  Rating    float64       `json:"rating"`
  Status    Status        `json:"status,omitempty"`
  CreatedAt time.Time     `json:"createdAt"`
}

type Reviews []Review
