package http_checker

import (
	"fmt"
	"net/http"
	"time"
)

func CheckHttpEndpoint(httpUrl string) bool {
	httpClient := http.Client{Timeout: time.Duration(10) * time.Second}

	resp, err := httpClient.Get(httpUrl)

	if err != nil {
		fmt.Printf("Error %s", err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}
