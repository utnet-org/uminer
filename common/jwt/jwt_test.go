package jwt

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("a22c85f6615e760cdd65d8ab27d06a67", "asdf", time.Second*time.Duration(86400))
	fmt.Println(token, err)
}
