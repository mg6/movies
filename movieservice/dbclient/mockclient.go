package dbclient

import (
  "github.com/stretchr/testify/mock"
  "github.com/mg6/movies/movieservice/model"
)

type MockClient struct {
  mock.Mock
}

func (m *MockClient) CreateMovie(movie *model.Movie) error {
  args := m.Mock.Called(movie)
  return args.Error(0)
}

func (m *MockClient) GetMovies() (model.Movies, error) {
  args := m.Mock.Called()
  return args.Get(0).(model.Movies), args.Error(1)
}

func (m *MockClient) DeleteMovie(id string) error {
  args := m.Mock.Called(id)
  return args.Error(0)
}