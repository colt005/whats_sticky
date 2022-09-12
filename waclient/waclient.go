package waclient

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
		return
	}
	fmt.Println(newFile.Name())

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
