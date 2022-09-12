package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleWhatsAppWebhook(c echo.Context) (err error) {

	json_map := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(json_map)

	return c.NoContent(200)
}

func GetHome(c echo.Context) (err error) {

	return c.String(http.StatusOK, "Simple WhatsApp Webhook tester</br>There is no front-end, see server.js for implementation!")
}

func GetHealth(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello world",
	})
}
