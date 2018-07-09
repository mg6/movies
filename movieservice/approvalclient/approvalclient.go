package approvalclient

import (
  . "github.com/mg6/movies/movieservice/model"
)

type ApprovalArgs struct {
  Review Review
}

type ApprovalReply struct {
  Status Status
}

type ApprovalClient interface {
  RequestApproval(review Review) (ApprovalReply, error)
}
