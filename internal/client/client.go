package client

import (
	"net/http"
	"io"
	"os"
	"fmt"
)

func GetRequest(path string) (io.ReadCloser, error) {

	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", os.Getenv("RV_API_URL"), path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", os.Getenv("RV_API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func PostRequest(path string, body io.Reader) (io.ReadCloser, error) {

	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", os.Getenv("RV_API_URL"), path)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", os.Getenv("RV_API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// defer client?

	return res.Body, nil
}