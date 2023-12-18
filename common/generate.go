package main

import (
	"uminer/common/utils"
)

//go:generate go run generate.go
func main() {
	err := utils.Generate()
	if err != nil {
		panic(err)
	}
}
