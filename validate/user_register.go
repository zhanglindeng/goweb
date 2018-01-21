package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"gopkg.in/go-playground/validator.v8"
)

type UserRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

func UserRegisterValidate(ctx *gin.Context) (*UserRegister, error) {

	decoder := json.NewDecoder(ctx.Request.Body)

	ur := &UserRegister{}
	if err := decoder.Decode(ur); err != nil {
		return ur, err
	}

	config := &validator.Config{TagName: "validate"}

	v := validator.New(config)

	if err := v.Struct(ur); err != nil {
		return ur, err
	}

	return ur, nil
}
