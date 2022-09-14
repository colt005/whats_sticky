package router

import (
	"github.com/colt005/whats_sticky/handler"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/", handler.GetHome)
	app.GET("/tmpfile", handler.GetTmpFile)
	app.GET("/health", handler.GetHealth)
	app.POST("/webhook", handler.HandleWhatsAppWebhook)
	app.GET("/webhook", handler.HandleWhatsAppWebhookVerify)
}
