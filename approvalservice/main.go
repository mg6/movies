package main

import (
  "log"
  "net/http"
  "github.com/gorilla/rpc"
  "github.com/gorilla/rpc/json"
  . "github.com/mg6/movies/approvalservice/service"
)

func main() {
  appName := "approvalservice"
  log.Printf("Starting %v\n", appName)

  s := rpc.NewServer()
  s.RegisterCodec(json.NewCodec(), "application/json")
  s.RegisterService(new(ApprovalService), "")
  http.Handle("/rpc", s)

  log.Println("Serving ...")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
