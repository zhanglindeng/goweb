package model

const (
	SubmenuStatusDisabled = iota
	SubmenuStatusEnable
)

type Submenu struct {
	BaseModel
	Name   string `json:"name" gorm:"type:varchar(50);not null;default:''"`
	Link   string `json:"link" gorm:"not null;default:''"`
	Sort   int    `json:"sort" gorm:"not null;default:0"`
	Status int    `json:"status" gorm:"not null;default:1"`
	MenuID int    `json:"menu_id" gorm:"index;not null;default:0"`
	Menu   Menu   `json:"menu"`
}
