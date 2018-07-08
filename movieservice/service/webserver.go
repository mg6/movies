package service

import (
  "net/http"
  "log"
  "github.com/mg6/movies/movieservice/dbclient"
)

func StartWebServer(port string) {
  r := NewRouter()
  http.Handle("/", r)

  log.Println("Staring HTTP service at " + port)
  err := http.ListenAndServe(":" + port, nil)

  if err != nil {
    log.Println("An error occurred starting HTTP server at port " + port)
    log.Println("Error: " + err.Error())
  }
}

func ConnectToDatabase(url string) {
  DbClient = &dbclient.MongoClient{}
  DbClient.Connect(url)
}
