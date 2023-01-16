package test

import (
	"github.com/shishanksingh2015/email-sample/config"
	"github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"testing"
)

func TestReadConfig(t *testing.T) {
	convey.Convey("Reading Config", t, func() {
		serverConfig := config.ReadConfig("..")
		convey.So(serverConfig.AuthSecret, convey.ShouldNotBeEmpty)
		convey.So(viper.GetString("authSecret"), convey.ShouldEqual, serverConfig.AuthSecret)
		convey.So(viper.GetString("sendGridApiKey"), convey.ShouldEqual, serverConfig.SendGridApiKey)
	})
}
