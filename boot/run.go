package boot

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/shishanksingh2015/email-sample/generatetoken"
	"github.com/shishanksingh2015/email-sample/health"
	"github.com/shishanksingh2015/email-sample/model"
	"github.com/shishanksingh2015/email-sample/sendmail"
)

func Run(serverConfig model.ServerConfig) *echo.Echo {

	server := echo.New()

	server.Logger.SetLevel(log.INFO)

	server.GET("/health", health.Get)

	tokenHandler := generatetoken.TokenHandler{Secret: serverConfig.AuthSecret}
	server.POST("/generate-token", tokenHandler.GenerateToken)

	//v1
	mailerClient := sendmail.New(serverConfig.SendGridApiKey)
	h := &sendmail.Handler{Mailer: mailerClient}
	v1 := server.Group("/v1")
	v1.POST("/sendmail", h.SendEmail)

	return server
}
