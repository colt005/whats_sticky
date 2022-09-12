package waclient

import (
	"fmt"
	"io"
	"net/http"
)

var httpClient *HTTPClient

func init() {
	httpClient = NewClient()
}

func GetMediaUrl(mediaId string) {
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

	fmt.Println(string(bodyBytes))
}


