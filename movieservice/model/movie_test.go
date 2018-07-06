package model

import (
  "sort"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestSortByRating(t *testing.T) {
  m1 := Movie{
    Rating: 2.0,
  }
  m2 := Movie{
    Rating: 0.0,
  }

  Convey("Given an unsorted movie list", t, func() {
    actual := Movies{m1, m2}
    So(actual[0].Rating, ShouldBeGreaterThan, actual[1].Rating)

    Convey("When movies are sorted by rating", func() {
      sort.Sort(ByRating(actual))

      Convey("Then the list should be sorted by rating", func() {
        expected := Movies{m2, m1}
        So(actual, ShouldResemble, expected)
      })
    })
  })
}
