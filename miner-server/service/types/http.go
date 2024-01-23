package types

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const mainURL = "https://console.utlab.io/openaiserver"

func HTTPRequest(method string, url string, data interface{}, contentType string, authToken string) []byte {
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)

	// 创建POST请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", authToken)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return result
}
