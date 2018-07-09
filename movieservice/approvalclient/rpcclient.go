package approvalclient

import (
  "bytes"
  "log"
  "net/http"
  "github.com/gorilla/rpc/json"
  . "github.com/mg6/movies/movieservice/model"
)

type ApprovalClientImpl struct {
  Url string
}

func (a *ApprovalClientImpl) RequestApproval(review Review) (ApprovalReply, error) {
  args := ApprovalArgs{Review: review}
  message, err := json.EncodeClientRequest("ApprovalService.GetApproval", args)
  if err != nil {
    log.Fatal(err)
  }

  req, err := http.NewRequest("POST", a.Url, bytes.NewBuffer(message))
  if err != nil {
    log.Fatal(err)
  }

  req.Header.Set("Content-Type", "application/json")
  client := new(http.Client)
  resp, err := client.Do(req)
  if err != nil {
    log.Fatalf("Error sending request to %s: %s", a.Url, err)
  }
  defer resp.Body.Close()

  var reply ApprovalReply
  err = json.DecodeClientResponse(resp.Body, &reply)
  if err != nil {
    log.Fatalf("Cannot decode response. %s", err)
  }

  return reply, nil
}
