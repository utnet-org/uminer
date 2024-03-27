package connect

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// MainURL container cloud server url
const MainURL = "https://console.utlab.io/openaiserver"

// Delay connection rpc delay
const Delay = 4

// HTTPRequest cloud container cloud server http request handler
func HTTPRequest(method string, url string, data interface{}, contentType string, authToken string) []byte {
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)

	// HTTP POST request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	// set header
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", "Bearer "+authToken)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read response contents
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return result
}
