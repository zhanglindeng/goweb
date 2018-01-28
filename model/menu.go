package model

const (
	MenuStatusDisabled = iota
	MenuStatusEnable
)

type Menu struct {
	BaseModel
	Name     string    `json:"name" gorm:"type:varchar(50);not null;default:''"`
	NameEn   string    `json:"name_en" gorm:"type:varchar(50);unique;not null;default:''"`
	Icon     string    `json:"icon"`
	Remark   string    `json:"remark"`
	Sort     int       `json:"sort" gorm:"not null;default:0"`
	Status   int       `json:"status" gorm:"not null;default:1"`
	Submenus []Submenu `json:"submenus"`
}
