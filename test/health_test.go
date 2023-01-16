package test

import (
	"github.com/shishanksingh2015/email-sample/boot"
	"github.com/shishanksingh2015/email-sample/config"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	Convey("Test health Endpoint ", t, func() {
		server := boot.Run(config.ReadConfig(".."))
		ts := httptest.NewServer(server)
		defer ts.Close()
		NewRequest := func(method, url string, body io.Reader) *http.Request {
			r, err := http.NewRequest(method, url, body)

			So(err, ShouldBeNil)

			if err != nil {
				t.Fatal(err)
			}
			return r
		}
		Convey("GET /health should return 200 OK\n", func() {
			req := NewRequest(http.MethodGet, ts.URL+"/health", nil)
			resp, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
		})
	})
}
