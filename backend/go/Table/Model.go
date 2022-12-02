package Table

import (
	Category "namazu/Category"

	_ "github.com/go-sql-driver/mysql"
)

type TableModel struct {
	Id              uint                   `json:"id"`
	Category        uint                   `gorm:"ForeignKey:ID" json:"category"`
	Code            int                    `json:"code"`
	Capacity        string                 `json:"capacity"`
	Reserved        bool                   `json:"reserved"`
	CategoryContent Category.CategoryModel `gorm:"ForeignKey:Category" json:"categorycontent"`
}

func (b *TableModel) TableName() string {
	return "tables"
}
