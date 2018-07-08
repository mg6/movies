package model

import (
  "time"
)

type Status string

const (
  Unapproved  Status = "unapproved"
  Approved    Status = "approved"
)

type Review struct {
  Text      string        `json:"text"`
  Rating    float64       `json:"rating"`
  Status    Status        `json:"status"`
  CreatedAt time.Time     `json:"createdAt"`
}

type Reviews []Review
