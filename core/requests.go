package core

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func makeRequest(service string, method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, service+path, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// MyRequest makes a request to a give url
func MyRequest(service string, method string, path string) ([]byte, int) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, _ := makeRequest(service, method, path, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, res.StatusCode
	}
	log.Printf("res.StatusCode %d", res.StatusCode)
	if res.StatusCode != 200 {
		return nil, res.StatusCode
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("001", err)
	}
	defer res.Body.Close()
	log.Printf("DATA %s", body)
	return body, res.StatusCode
}
