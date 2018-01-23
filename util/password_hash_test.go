package util

import "testing"

func TestPasswordHash(t *testing.T) {

	if hash, err := PasswordHash("123456"); err != nil {
		t.Error(err)
	} else {
		t.Log(hash)
	}
}
