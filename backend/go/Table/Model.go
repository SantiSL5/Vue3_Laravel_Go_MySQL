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

//GetAllTables Fetch all table data
func GetAllTablesDB(c *gin.Context) []TableModel {

	var tables []TableModel

	if err := Config.DB.Preload("CategoryContent").Find(&tables).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return tables

}

func GetOneTableDB(id int, c *gin.Context) (TableModel, error) {

	var table TableModel

	err := Config.DB.Preload("CategoryContent").Where("ID = ?", id).Find(&table).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return table, err
}
