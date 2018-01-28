package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"gopkg.in/go-playground/validator.v8"
)

type MenuAdd struct {
	Name   string `json:"name" validate:"required,max=2"`
	NameEn string `json:"name_en" validate:"required,max=20"`
	Icon   string `json:"icon" validate:"required"`
	Sort   int    `json:"sort" validate:"required"`
	Remark string `json:"remark"`
}

var MenuUpdateValidate = MenuAddValidate

func MenuAddValidate(ctx *gin.Context) (*MenuAdd, error) {

	decoder := json.NewDecoder(ctx.Request.Body)

	ma := &MenuAdd{}
	if err := decoder.Decode(ma); err != nil {
		return ma, err
	}

	config := &validator.Config{TagName: "validate"}

	v := validator.New(config)

	if err := v.Struct(ma); err != nil {
		return ma, err
	}

	return ma, nil
}

