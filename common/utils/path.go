package utils

import (
	"regexp"
)

func OptimizeSeparator(path string) string {
	reg := regexp.MustCompile("///*")
	return reg.ReplaceAllString(path, "/")
}
