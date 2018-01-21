package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"gopkg.in/go-playground/validator.v8"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func UserLoginValidate(ctx *gin.Context) (*UserLogin, error) {

	decoder := json.NewDecoder(ctx.Request.Body)

	ul := &UserLogin{}
	if err := decoder.Decode(ul); err != nil {
		return ul, err
	}

	config := &validator.Config{TagName: "validate"}

	v := validator.New(config)

	if err := v.Struct(ul); err != nil {
		return ul, err
	}

	return ul, nil
}
