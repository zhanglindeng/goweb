package util

import "testing"

func TestSha1(t *testing.T) {
	if Sha1("123456") != "7c4a8d09ca3762af61e59520943dc26494f8941b" {
		t.Error("Sha1")
	}
}
