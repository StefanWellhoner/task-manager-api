package model

type Category struct {
	Base
	Name  string
	Tasks []Task `gorm:"foreignKey:CategoryID"`
}
