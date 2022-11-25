package Table

import (
	"fmt"
	"namazu/Config"

	_ "github.com/go-sql-driver/mysql"
)

type TableModel struct {
	Id      	uint   	`json:"id"`
	Code    	int 	`json:"code"`
	Category   	string 	`json:"category"`
	Capacity   	string 	`json:"capacity"`
	Reserved 	bool 	`json:"reserved"`
}

func (b *TableModel) TableName() string {
	return "tables"
}

//GetAllTables Fetch all table data
func GetAllTables(able *[]Table) (err error) {
	if err = Config.DB.Find(table).Error; err != nil {
		return err
	}
	return nil
}

//CreateTable ... Insert New data
func CreateTable(table *Table) (err error) {
	if err = Config.DB.Create(table).Error; err != nil {
		return err
	}
	return nil
}

//GetTableByID ... Fetch only one table by Id
func GetTableByID(table *Table, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(table).Error; err != nil {
		return err
	}
	return nil
}

//UpdateTable ... Update table
func UpdateTable(table *Table, id string) (err error) {
	fmt.Println(table)
	Config.DB.Save(table)
	return nil
}

//DeleteTable ... Delete table
func DeleteTable(table *Table, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(table)
	return nil
}
