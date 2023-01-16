package test

import (
	"bytes"
	"encoding/json"
	"github.com/shishanksingh2015/email-sample/boot"
	"github.com/shishanksingh2015/email-sample/config"
	"github.com/shishanksingh2015/email-sample/generatetoken"
	"github.com/shishanksingh2015/email-sample/model"
	"github.com/shishanksingh2015/email-sample/sendmail"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

func GetToken() string {
	request := generatetoken.Claims{
		RequestorEmail: "singhshishank2012@gmail.com",
		TraceId:        "xxxxx-xxxx-xxxxx-xxxx",
	}
	claims, _ := json.Marshal(request)
	token, _ := generatetoken.CreateToken(string(claims), viper.GetString("authSecret"))
	return token
}

func TestSendEmail(t *testing.T) {
	Convey("Test SendMail Endpoint ", t, func() {
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
		Convey("It should not return error when valid requests is sent", func() {
			var validReqData = `{"from": "singhshishank2012@gmail.com","from_name": "senderName","to": "singhshishank2012@gmail.com",
				"to_name": "recipientName","subject": "demo for api test","content": "demo tested"}`
			req := NewRequest(http.MethodPost, ts.URL+"/v1/sendmail?", bytes.NewBuffer([]byte(validReqData)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+GetToken())
			resp, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			apiResponse := &model.ApiStatusResponse{}
			err = json.NewDecoder(resp.Body).Decode(apiResponse)
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			So(apiResponse.Message, ShouldEqual, sendmail.EmailSent)
			So(apiResponse.TraceId, ShouldNotBeBlank)
		})
		Convey("It should return error when required fields are wrong type", func() {
			var inValidReqData = `{"from": 1,"from_name": "senderName","to": "shishank.com",
				"to_name": "recipientName","subject": "demo for api test","content": "demo tested"}`

			req := NewRequest(http.MethodPost, ts.URL+"/v1/sendmail?", bytes.NewBuffer([]byte(inValidReqData)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+GetToken())
			resp, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			apiResponse := &model.ApiStatusResponse{}
			err = json.NewDecoder(resp.Body).Decode(apiResponse)
			So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(apiResponse.Message, ShouldEqual, sendmail.TypeError)
			So(apiResponse.TraceId, ShouldNotBeBlank)
		})
		Convey("It should return error when required fields are invalid", func() {
			var inValidReqData = `{
				"from": "shishank",
				"from_name": "senderName",
				"to": "singhshishank2012@gmail.com",
				"to_name": "recipientName",
				"subject": "demo for api test",
				"content": "demo tested"
			}`

			req := NewRequest(http.MethodPost, ts.URL+"/v1/sendmail?", bytes.NewBuffer([]byte(inValidReqData)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+GetToken())
			resp, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			apiResponse := &model.ApiStatusResponse{}
			err = json.NewDecoder(resp.Body).Decode(apiResponse)
			So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(apiResponse.Message, ShouldEqual, sendmail.ValidationError)
			So(apiResponse.TraceId, ShouldNotBeBlank)
		})
		Convey("It should return error when required fields are missing", func() {
			var inValidReqData = `{
				"from": shishank,
				"from_name": "senderName",
				"to_name": "recipientName",
				"subject": "demo for api test",
				"content": "demo tested"
			}`

			req := NewRequest(http.MethodPost, ts.URL+"/v1/sendmail?", bytes.NewBuffer([]byte(inValidReqData)))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+GetToken())
			resp, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			apiResponse := &model.ApiStatusResponse{}
			err = json.NewDecoder(resp.Body).Decode(apiResponse)
			So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(apiResponse.Message, ShouldEqual, sendmail.TypeError)
			So(apiResponse.TraceId, ShouldNotBeBlank)
		})
	})
}
