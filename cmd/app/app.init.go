package main

import (
	"fmt"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/routes"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/config"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"log"
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
	db, err := config.PostgresInit(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	// initiate validator
	validate := validator.New()

	// set log level
	gin.SetMode(cfg.App.LogLevel)

	// create new gin server
	app := gin.New()

	app.Use(utils.SecurityHeaders())
	app.Use(utils.CorsConfig())

	appUserRoute := routes.NewRoute(db, app, *cfg, validate)
	appUserRoute.RouteInit()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: app.Handler(),
	}

	log.Println("Server is running at ", srv.Addr)
	return srv
}
