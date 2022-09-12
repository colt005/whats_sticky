package handler

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

func HandleWhatsAppWebhook(c echo.Context) (err error) {

	json_map := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(json_map)

	return
}
