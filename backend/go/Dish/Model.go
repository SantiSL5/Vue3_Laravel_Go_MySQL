package Dish

import (
	DishType "namazu/DishType"

	_ "github.com/go-sql-driver/mysql"
)

type DishModel struct {
	Id          uint                   `json:"id"`
	Dish        string                 `json:"dish"`
	Type        uint                   `gorm:"ForeignKey:ID" json:"type"`
	Price       float64                `json:"price"`
	Desc        string                 `json:"desc"`
	Photo       string                 `json:"photo"`
	TypeContent DishType.DishTypeModel `gorm:"ForeignKey:Type" json:"typecontent"`
}

func (b *DishModel) TableName() string {
	return "dishes"
}
