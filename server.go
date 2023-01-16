package main

import (
	"github.com/labstack/gommon/log"
	"github.com/shishanksingh2015/email-sample/boot"
	"github.com/shishanksingh2015/email-sample/config"
	"net/http"
	"time"
)

func main() {

	serverConfig := config.ReadConfig(".")
	server := boot.Run(serverConfig)
	s := &http.Server{
		Addr:         ":5000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.StartServer(s))
}
