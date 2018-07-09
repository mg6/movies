package service

import (
  "net/http"
  "log"
  "github.com/mg6/movies/movieservice/dbclient"
  "github.com/mg6/movies/movieservice/approvalclient"
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

func InitializeApprovalClient(url string) {
  log.Printf("Configuring approval service at %v\n", url)
  ApprovalClient = &approvalclient.ApprovalClientImpl{Url: url}
}
