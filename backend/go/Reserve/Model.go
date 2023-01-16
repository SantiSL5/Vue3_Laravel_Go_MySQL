package Reserve

import (
	User   "namazu/User"
	Table  "namazu/Table"

	_ "github.com/go-sql-driver/mysql"
)

type ReserveModel struct {
	Id              uint                   `json:"id"`
	User            uint                   `gorm:"ForeignKey:ID" json:"user"`
	Table           uint                   `gorm:"ForeignKey:ID" json:"table"`
	Date            string                 `json:"date"`
	Turn            string                 `json:"turn"`
	Confirmed       bool                   `json:"confirmed"`
	UserContent     User.UserModel         `gorm:"ForeignKey:User"  json:"usercontent"`
	TableContent    Table.TableModel       `gorm:"ForeignKey:Table" json:"tablecontent"`
}

func (b *ReserveModel) TableName() string {
	return "reserves"
}
