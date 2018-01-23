package test

import (
	"encoding/json"
	"testing"

	"github.com/parnurzeal/gorequest"
)

func TestUserLogin(t *testing.T) {

	resp, body, errs := gorequest.New().Post("http://localhost:10086/api/user/login").
		Set("Content-Type", "application/json").
		Send(`{"email":"zhangsan@163.com","password":"123456"}`).
		End()

	if len(errs) > 0 {
		t.Error(errs)
		return
	}

	if resp.StatusCode != 200 {
		t.Error("status code", resp.StatusCode, resp.Status)
		return
	}

	var data struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Token   string `json:"token"`
	}

	if err := json.Unmarshal([]byte(body), &data); err != nil {
		t.Error(err)
		return
	}

	if data.Code != 0 {
		t.Error("code", data.Code, data.Message)
		return
	}

	t.Log(data.Token)
}
