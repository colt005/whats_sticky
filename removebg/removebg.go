package removebg

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type HTTPClient struct {
	client *http.Client
}

var httpClient *HTTPClient

func init() {
	httpClient = NewClient()
}

//Do dispatches the HTTP request to the network
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func NewClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

func GetSticker(localPath string) (localFilePath string) {
	file, err := os.Open(localPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", "https://removeg.kudla.live/api/remove-bg", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.client.Do(r)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fName := uuid.New().String()
	fileName := fName + ".webp"
	tmpPath := filepath.Join("/tmp", fileName)

	newFile, err := os.Create(tmpPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer newFile.Close()

	if _, err = newFile.Write(b); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(newFile.Name())

	localFilePath = tmpPath

	return
}

func GetStickerFromVideo(localPath string) (localFilePath string) {
	file, err := os.Open(localPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", "https://removeg.kudla.live/api/mp4-webp", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.client.Do(r)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fName := uuid.New().String()
	fileName := fName + ".webp"
	tmpPath := filepath.Join("/tmp", fileName)

	newFile, err := os.Create(tmpPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer newFile.Close()

	if _, err = newFile.Write(b); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(newFile.Name())

	localFilePath = tmpPath

	return
}
