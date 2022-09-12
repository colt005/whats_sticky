package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/colt005/whats_sticky/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/logger"
)

func HandleWhatsAppWebhookVerify(c echo.Context) (err error) {

	json_map := make(map[string]interface{})
	headerChallenge := c.QueryParams().Get("hub.challenge")
	verifyToken := c.QueryParams().Get("hub.verify_token")

	err = json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(json_map)

	if verifyToken == config.Config("VERIFY_TOKEN") {
		return c.String(http.StatusOK, headerChallenge)
	} else {
		return c.String(http.StatusUnauthorized, "lol nice try")
	}
}

func HandleWhatsAppWebhook(c echo.Context) (err error) {

	json_map := make(map[string]interface{})
	headerChallenge := c.QueryParams().Get("hub.challenge")

	err = json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(json_map)

	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		logger.Error(err)
	}

	fmt.Println(string(bodyBytes))

	return c.String(http.StatusOK, headerChallenge)
}

func GetHome(c echo.Context) (err error) {

	return c.String(http.StatusOK, "Simple WhatsApp Webhook tester</br>There is no front-end, see server.js for implementation!")
}

func GetHealth(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello world",
	})
}
