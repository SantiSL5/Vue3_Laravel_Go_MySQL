package Table

import (
	"fmt"
	Category "namazu/Category"
	"namazu/Config"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type TableModel struct {
	Id       uint                   `json:"id"`
	Code     int                    `json:"code"`
	Category Category.CategoryModel `gorm:"ForeignKey:Id" json:"category"`
	Capacity string                 `json:"capacity"`
	Reserved bool                   `json:"reserved"`
}

func (b *TableModel) TableName() string {
	return "tables"
}

//GetAllTables Fetch all table data
func GetAllTablesDB(c *gin.Context) []TableModel {

	var tables []TableModel

	if err := Config.DB.Find(&tables).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return tables

}

func GetOneTableDB(id int, c *gin.Context) TableModel {

	var table TableModel

	if err := Config.DB.Where("ID = ?", id).Find(&table).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return table
	} else {
		return table
	}

}
