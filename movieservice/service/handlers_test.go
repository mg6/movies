package service

import (
  "net/http/httptest"
  "testing"
  . "github.com/smartystreets/goconvey/convey"
)

func TestGetWrongPath(t *testing.T) {
  Convey("Given a HTTP request for /invalid/123", t, func() {
    req := httptest.NewRequest("GET", "/invalid/123", nil)
    resp := httptest.NewRecorder()

    Convey("When the request is handled by the router", func() {
      NewRouter().ServeHTTP(resp, req)

      Convey("Then the response should be a 404", func() {
        So(resp.Code, ShouldEqual, 404)
      })
    })
  })
}
