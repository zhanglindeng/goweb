package validate

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"gopkg.in/go-playground/validator.v8"
)

type MenuAddSubmenu struct {
	Name string `json:"name" validate:"required,max=6"`
	Link string `json:"link" validate:"required"`
	Sort int    `json:"sort" validate:"required"`
	//MenuID int    `json:"menu_id" validate:"required,min=1"`
}

var MenuUpdateSubmenuValidate = MenuAddSubmenuValidate

func MenuAddSubmenuValidate(ctx *gin.Context) (*MenuAddSubmenu, error) {

	decoder := json.NewDecoder(ctx.Request.Body)

	mas := &MenuAddSubmenu{}
	if err := decoder.Decode(mas); err != nil {
		return mas, err
	}

	config := &validator.Config{TagName: "validate"}

	v := validator.New(config)

	if err := v.Struct(mas); err != nil {
		return mas, err
	}

	return mas, nil
}
