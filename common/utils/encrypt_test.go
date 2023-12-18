package utils

import "testing"

func TestEncryptPassword(t *testing.T) {
	p, err := EncryptPassword("hello world")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(p)
}
