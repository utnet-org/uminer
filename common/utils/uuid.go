package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetUUIDWithoutSeparator() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// uuid手字母替换为字符 k8s的service首字母不能为数字
func GetUUIDStartWithAlphabetic() string {
	uuid := GetUUIDWithoutSeparator()
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	uuid = string(bytes[r.Intn(len(bytes))]) + uuid[1:]
	return uuid
}
