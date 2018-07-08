package main

import (
  "fmt"
  "github.com/mg6/movies/movieservice/service"
  "github.com/namsral/flag"
)

func main() {
  appName := "movieservice"
  fmt.Printf("Starting %v\n", appName)

  dbDialUrl := flag.String("mongodb-url", "localhost", "MongoDB connection URL")
  flag.Parse()

  service.ConnectToDatabase(*dbDialUrl)
  service.StartWebServer("8080")
}
