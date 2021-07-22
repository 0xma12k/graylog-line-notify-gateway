package main

import (
	"net/http"

	"github.com/0xma12k/graylog-line-notify-gateway/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Init()
	logrus.Info("server is running on port " + config.Get().Port)
	http.ListenAndServe(":"+config.Get().Port, Router())
}
