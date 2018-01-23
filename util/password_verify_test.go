package util

import "testing"

func TestPasswordVerify(t *testing.T) {
	// 在 hash 最后添加内容一样可以通过，不知道是为什么
	if !PasswordVerify("123456", "$2a$10$7TtZkhDbKM3W9Gm23B50uuKQNi3oHOaHv/NhNeOBfdksQ567iS3Ye") {
		t.Error("PasswordVerify")
	}
}
