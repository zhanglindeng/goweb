package test

import (
	"encoding/json"
	"testing"

	"github.com/parnurzeal/gorequest"
)

func TestLogIndex(t *testing.T) {

	resp, body, errs := gorequest.New().Get("http://localhost:10086/api/log").
		Set("Content-Type", "application/json").
		Set("Authorization", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTY3Mjc2ODcsImlzcyI6InpoYW5nc2FuMkAxNjMuY29tIiwibmJmIjoxNTE2NzIwNDg3fQ.DjGoaqt_sWjZZ9upxLGdO3v4eoupfCHAPtlpZSvn-bc`).
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
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Logs    interface{} `json:"logs"`
	}

	if err := json.Unmarshal([]byte(body), &data); err != nil {
		t.Error(err)
		return
	}

	if data.Code != 0 {
		t.Error("code", data.Code, data.Message)
		return
	}

	t.Log(data.Logs)
}
