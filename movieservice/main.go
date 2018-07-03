package main

import (
  "fmt"
  "github.com/mg6/movies/movieservice/service"
)

func main() {
  appName := "movieservice"
  fmt.Printf("Starting %v\n", appName)
  service.StartWebServer("8080")
}
