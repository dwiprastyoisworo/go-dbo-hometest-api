package main

import (
	"fmt"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func AppInit() *http.Server {
	// setup logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	// setup app config
	cfg, err := config.AppConfigInit()
	if err != nil {
		panic(err)
	}

	// setup postgres connection
	_, err = config.PostgresInit(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	// set log level
	gin.SetMode(cfg.App.LogLevel)

	// create new gin server
	app := gin.New()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: app.Handler(),
	}
	return srv
}
