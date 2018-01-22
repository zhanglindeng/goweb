package util

import "testing"

func TestMd5(t *testing.T) {
	// e10adc3949ba59abbe56e057f20f883e
	t.Log(Md5("123456"))
}
