package main

import (
	"log"

	"github.com/colt005/whats_sticky/config"
	"github.com/colt005/whats_sticky/database"
	router "github.com/colt005/whats_sticky/router"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	app := echo.New()
	err := database.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	router.SetupRoutes(app)
	logrus.Fatal(app.Start(":" + config.Config("PORT")))

}
