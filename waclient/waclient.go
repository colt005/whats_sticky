package waclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/colt005/whats_sticky/config"
	"github.com/colt005/whats_sticky/models"
	"github.com/google/uuid"
)

type HTTPClient struct {
	client *http.Client
}

func init() {
	httpClient = NewClient()
}

func GetMediaUrl(mediaId string) (mediaResponse *models.MediaResponse, err error) {
	req, err := http.NewRequest("GET", "https://graph.facebook.com/v13.0/"+mediaId, nil)

	if err != nil {
		fmt.Println(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		mR, err := models.UnmarshalMediaResponse(bodyBytes)
		if err != nil {
			return nil, err
		}

		mediaResponse = &mR

		return mediaResponse, nil
	} else {
		fmt.Println(string(bodyBytes))
		return nil, errors.New("failed to get media")

	}

}

func DownloadMedia(mediaResponse models.MediaResponse) (localPath string, err error) {

	req, err := http.NewRequest("GET", mediaResponse.URL, nil)

	if err != nil {
		fmt.Println(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode == http.StatusOK {

		fName := uuid.New().String()
		fileName := fName + getFileExtension(mediaResponse.MIMEType)
		tmpPath := filepath.Join("/tmp", fileName)
		if getFileExtension(mediaResponse.MIMEType) == "" {
			fmt.Println("Failed to get file extension")
			return
		}
		newFile, err := os.Create(tmpPath)
		if err != nil {
			fmt.Println(err.Error())
		}

		defer newFile.Close()

		if _, err = newFile.Write(bodyBytes); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(newFile.Name())
		localPath = tmpPath
	} else {
		fmt.Println(string(bodyBytes))
	}

	return
}

func SendTextMessage(mobileNo string, message string) {
	url := "https://graph.facebook.com/v13.0/" + config.Config("MOBILE_ID") + "/messages"
	method := "POST"

	t := models.TextMessageRequest{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               mobileNo,
		Type:             "text",
		Text: models.Text{
			PreviewURL: false,
			Body:       message,
		},
	}

	reqBody, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(reqBody))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bodyBytes))

}

func SendStickerById(stickerId string, mobileNo string) (err error) {

	s := &models.StickerRequest{}

	s.MessagingProduct = "whatsapp"
	s.RecipientType = "individual"
	s.Type = "sticker"
	s.To = mobileNo
	s.Sticker.ID = stickerId

	reqBody, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", "https://graph.facebook.com/v13.0/"+config.Config("MOBILE_ID")+"/messages", bytes.NewBuffer(reqBody))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		return err
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("asdasdsad")
	fmt.Println(string(bodyBytes))

	defer resp.Body.Close()

	return

}

func CreateImageFormFile(w *multipart.Writer, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	w.WriteField("messaging_product", "whatsapp")
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	h.Set("Content-Type", "image/webp")
	return w.CreatePart(h)
}

func UploadSticker(webpPath string) (mediaId string, err error) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("messaging_product", "whatsapp")
	file, errFile2 := os.Open(webpPath)
	if errFile2 != nil {
		fmt.Println(err)
	}
	defer file.Close()
	part2, errFile2 := CreateImageFormFile(writer, webpPath)
	if errFile2 != nil {
		fmt.Println(errFile2)
		return
	}
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2)
		return
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	r, _ := http.NewRequest("POST", "https://graph.facebook.com/v14.0/"+config.Config("MOBILE_ID")+"/media", payload)
	r.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.Do(r)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	if resp.StatusCode == http.StatusOK {
		var result map[string]string

		err = json.Unmarshal(b, &result)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result["id"])
		mediaId = result["id"]
	}

	return
}

func getFileExtension(mimeType string) (extension string) {
	switch mimeType {
	case "image/jpeg", "image/jpg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/webp":
		return ".webp"

	}
	return
}
