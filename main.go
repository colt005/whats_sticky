package main

import (
	"github.com/colt005/whats_sticky/config"
	router "github.com/colt005/whats_sticky/router"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	app := echo.New()

	router.SetupRoutes(app)
	logrus.Fatal(app.Start(":" + config.Config("PORT")))

}
