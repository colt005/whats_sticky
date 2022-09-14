package router

import (
	"github.com/labstack/echo/v4"
	"github.com/colt005/whats_sticky/stickyhandler"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/", stickyhandler.GetHome)
	app.GET("/tmpfile", stickyhandler.GetTmpFile)
	app.GET("/health", stickyhandler.GetHealth)
	app.POST("/webhook", stickyhandler.HandleWhatsAppWebhook)
	app.GET("/webhook", stickyhandler.HandleWhatsAppWebhookVerify)
}
