package model

import (
  "sort"
  "strings"
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

func TestIsValid(t *testing.T) {
  Convey("Given a movie with valid title of <Title>", t, func() {
    movie := Movie{
      Title: "Title",
    }

    Convey("When the movie is validated", func() {
      actual := movie.IsValid()

      Convey("Then the result should be successful", func() {
        expected := true
        So(actual, ShouldEqual, expected)
      })
    })
  })

  for _, title := range []string{"", "Ab", strings.Repeat("x", 51), "Test123"} {
    Convey("Given a movie with invalid title of <" + title + ">", t, func() {
      movie := Movie{
        Title: title,
      }

      Convey("When the movie is validated", func() {
        actual := movie.IsValid()

        Convey("Then the result should be unsuccessful", func() {
          expected := false
          So(actual, ShouldEqual, expected)
        })
      })
    })
  }
}
