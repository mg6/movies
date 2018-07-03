package model

type Movie struct {
  Id        string    `json:"id"`
  Title     string    `json:"title"`
  Rating    float64   `json:"rating"`
  Director  string    `json:"director"`
  Actors    []string  `json:"actors"`
}

type Movies []Movie
