package router

import (
	handler "github.com/colt005/whats_sticky/handler"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/health",handler.GetHealth)
	app.POST("/webhook", handler.HandleWhatsAppWebhook)
}
