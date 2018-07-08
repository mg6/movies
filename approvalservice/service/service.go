package approvalservice

import (
  "net/http"
  . "github.com/mg6/movies/approvalservice/model"
)

type ApprovalArgs struct {
  Review Review
}

type ApprovalReply struct {
  Status Status
}

type ApprovalService struct{}

func (s *ApprovalService) GetApproval(r *http.Request, args *ApprovalArgs, reply *ApprovalReply) error {
  reply.Status = Approved
  return nil
}
