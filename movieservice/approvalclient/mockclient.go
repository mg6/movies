package approvalclient

import (
  "github.com/stretchr/testify/mock"
  . "github.com/mg6/movies/movieservice/model"
)

type MockClient struct {
  mock.Mock
}

func (m *MockClient) RequestApproval(review Review) (ApprovalReply, error) {
  args := m.Mock.Called(review)
  return args.Get(0).(ApprovalReply), args.Error(1)
}
