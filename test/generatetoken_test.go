package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/shishanksingh2015/email-sample/boot"
	"github.com/shishanksingh2015/email-sample/config"
	"github.com/shishanksingh2015/email-sample/generatetoken"
	"github.com/shishanksingh2015/email-sample/model"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAndValidateToken(t *testing.T) {
	request := generatetoken.Claims{
		RequestorEmail: "singhshishank2012@gmail.com",
		TraceId:        "xxxxx-xxxx-xxxxx-xxxx",
	}
	claims, _ := json.Marshal(request)
	token, err := generatetoken.CreateToken(string(claims), viper.GetString("authSecret"))
	assert.NoError(t, err)
	c, err := GetClaims(token)
	assert.NoError(t, err)
	info := c["info"]
	assert.Equal(t, `{"requestorEmail":"singhshishank2012@gmail.com","traceId":"xxxxx-xxxx-xxxxx-xxxx"}`, info)
}

func GetClaims(encoded string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(encoded,
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("token not valid")
			}
			return []byte(viper.GetString("authSecret")), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")

	}
	return claims, nil
}

func TestGenerateToken(t *testing.T) {
	Convey("Test Generate token endpoint ", t, func() {
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
		Convey("POST /generate-token should return 200 OK", func() {
			var validReqData = `{"requestorEmail": "singhshishank2012@gmail.com"}`
			req := NewRequest(http.MethodPost, ts.URL+"/generate-token", bytes.NewBuffer([]byte(validReqData)))
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
		})
		Convey("POST /generate-token should return error for wrong type field", func() {
			var validReqData = `{"requestorEmail": 1}`
			req := NewRequest(http.MethodPost, ts.URL+"/generate-token", bytes.NewBuffer([]byte(validReqData)))
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			apiResponse := &model.ApiStatusResponse{}
			err = json.NewDecoder(resp.Body).Decode(apiResponse)
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(apiResponse.Message, ShouldEqual, generatetoken.TypeError)
		})
		Convey("POST /generate-token should return error for invalid field", func() {
			var validReqData = `{"requestorEmail": "shishank"}`
			req := NewRequest(http.MethodPost, ts.URL+"/generate-token", bytes.NewBuffer([]byte(validReqData)))
			req.Header.Set("Content-Type", "application/json")
			resp, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			apiResponse := &model.ApiStatusResponse{}
			err = json.NewDecoder(resp.Body).Decode(apiResponse)
			So(err, ShouldBeNil)
			So(resp.StatusCode, ShouldEqual, http.StatusBadRequest)
			So(apiResponse.Message, ShouldEqual, generatetoken.ValidationError)
		})
	})
}
