package model

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"type:varchar(128);unique"`
	Password string `json:"password"`
}
