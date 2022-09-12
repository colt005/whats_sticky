package router

import (
	handler "github.com/colt005/whats_sticky/handler"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.POST("/whatsapp/webhook", handler.HandleWhatsAppWebhook)
}
