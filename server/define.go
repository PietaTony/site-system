package server

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	port = ":8080"
)

const (
	layout = "2006-01-02 15:04:05"
)

const beginTimeout = 5

func getJSONAPI(url string, target interface{}) error {
	client := http.Client{
		Timeout: 50 * time.Millisecond,
	}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
