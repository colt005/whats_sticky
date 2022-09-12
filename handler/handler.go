package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleWhatsAppWebhook(c echo.Context) (err error) {

	json_map := make(map[string]interface{})
	headerChallenge := c.Request().Header.Get("hub.challenge")
	formChallenge := c.Request().Form.Get("hub.challenge")
	postFormChallenge := c.Request().PostForm.Get("hub.challenge")
	fmt.Println("Header Challenge")
	fmt.Println(headerChallenge)
	fmt.Println("Form Challenge")
	fmt.Println(formChallenge)
	fmt.Println("Post Form Challenge")
	fmt.Println(postFormChallenge)
	fmt.Println(c.Request().Body)
	fmt.Println(c.Request().Header)
	fmt.Println(c.QueryParams())
	err = json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(json_map)

	return c.String(http.StatusOK, "hellotoken")
}

func GetHome(c echo.Context) (err error) {

	return c.String(http.StatusOK, "Simple WhatsApp Webhook tester</br>There is no front-end, see server.js for implementation!")
}

func GetHealth(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello world",
	})
}
