package stickyhandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/colt005/whats_sticky/config"
	"github.com/colt005/whats_sticky/models"
	"github.com/colt005/whats_sticky/removebg"
	"github.com/colt005/whats_sticky/waclient"
	"github.com/labstack/echo/v4"
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

	// headerChallenge := c.QueryParams().Get("hub.challenge")

	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bodyBytes))

	messageResponse, err := models.UnmarshalMessageResponse(bodyBytes)

	if err != nil {
		fmt.Println(err)
	}
	var mediaResponse *models.MediaResponse
	// if len(messageResponse.Entry) > 0 && len(messageResponse.Entry[0].Changes) > 0 && len(messageResponse.Entry[0].Changes[0].Value.Messages) > 0 && messageResponse.Entry[0].Changes[0].Value.Messages[0].Type == "image" {
	// 	mediaResponse, err = waclient.GetMediaUrl(messageResponse.Entry[0].Changes[0].Value.Messages[0].Image.ID)
	// } else {
	// 	return
	// }
	for _, e := range messageResponse.Entry {
		for _, c2 := range e.Changes {
			for _, m := range c2.Value.Messages {
				contact, contactErr := waclient.GetFirstContact(c2.Value.Contacts)
				if contactErr != nil {
					fmt.Println(err)
					if err.Error() == waclient.NO_PROFILE_NAME {
						contact = &models.Contact{
							Profile: models.Profile{
								Name: "User",
							},
						}
					}
				}
				if m.Type == "image" {
					fmt.Println(m.ID)
					fmt.Println(m)
					go waclient.MarkMessageAsRead(m.ID)
					waclient.SendTextMessage(m.From, fmt.Sprintf("Hi %s! \nPlease wait while I get my hands sticky and work on your sticker!", contact.Profile.Name))
					mediaResponse, err = waclient.GetMediaUrl(m.Image.ID)

					var filesToRemove []string

					localPath, err := waclient.DownloadMedia(*mediaResponse)
					fmt.Println(localPath)
					filesToRemove = append(filesToRemove, localPath)
					if err != nil {
						fmt.Println(err)
					}

					stickerPath := removebg.GetSticker(localPath)
					filesToRemove = append(filesToRemove, stickerPath)

					fmt.Println(stickerPath)

					mediaId, err := waclient.UploadSticker(stickerPath)

					if err != nil {
						fmt.Println(err)
					}

					err = waclient.SendStickerById(mediaId, m.From)
					if err != nil {
						fmt.Println(err)
					}

					for _, v := range filesToRemove {
						os.Remove(v)
					}
				} else if m.Type == "video" {
					go waclient.MarkMessageAsRead(m.ID)
					waclient.SendTextMessage(m.From, fmt.Sprintf("Hi %s! \nPlease wait while I get my hands sticky and work on your sticker!", contact.Profile.Name))
					mediaResponse, err = waclient.GetMediaUrl(m.Video.ID)

					var filesToRemove []string

					localPath, err := waclient.DownloadMedia(*mediaResponse)
					fmt.Println(localPath)
					filesToRemove = append(filesToRemove, localPath)
					if err != nil {
						fmt.Println(err)
					}

					stickerPath := removebg.GetStickerFromVideo(localPath)
					filesToRemove = append(filesToRemove, stickerPath)

					fmt.Println(stickerPath)

					mediaId, err := waclient.UploadSticker(stickerPath)

					if err != nil {
						fmt.Println(err)
					}

					err = waclient.SendStickerById(mediaId, m.From)
					if err != nil {
						fmt.Println(err)
					}

					for _, v := range filesToRemove {
						os.Remove(v)
					}

				} else if m.Type == "text" {
					fmt.Println("Got a message")
					fmt.Println(m.ID)
					fmt.Println(m)
					go waclient.MarkMessageAsRead(m.ID)
					waclient.SendTextMessage(m.From, fmt.Sprintf("Hi %s! \nShoot me an image and I will reply with a sticker!", contact.Profile.Name))
				}
			}
		}
	}

	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusOK)
}

func GetHome(c echo.Context) (err error) {

	return c.String(http.StatusOK, "Simple WhatsApp Webhook tester</br>There is no front-end, see server.js for implementation!")
}

func GetTmpFile(c echo.Context) (err error) {

	var files []string

	root := "/tmp"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Split(info.Name(), ".")[len(strings.Split(info.Name(), "."))-1] == "jpg" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	return c.Attachment(files[0], "adad.jpg")
}

func GetHealth(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, echo.Map{
		"message": "hello world",
	})
}
