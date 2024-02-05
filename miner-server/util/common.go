package util

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

const alphanumerics = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 用 Gzip 压缩了，因此网关需判断头里是否有此 header 头，再用 compress/gzip 解压即可。
func GzipApi(res *http.Response) []byte {
	// 是否有 gzip
	gzipFlag := false
	for k, v := range res.Header {
		if strings.ToLower(k) == "content-encoding" && strings.ToLower(v[0]) == "gzip" {
			gzipFlag = true
		}
	}
	var content []byte
	if gzipFlag {
		// 创建 gzip.Reader
		gr, err := gzip.NewReader(res.Body)
		defer gr.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		content, _ = ioutil.ReadAll(gr)
	} else {
		content, _ = ioutil.ReadAll(res.Body)
	}
	return content
}

// RandomString generates a pseudo-random string of length n.
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanumerics[rand.Int63()%int64(len(alphanumerics))]
	}
	return string(b)
}
