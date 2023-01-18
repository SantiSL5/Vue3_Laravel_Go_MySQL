package Table

import (
	"fmt"
	"namazu/Config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOneTableRepo(id int, c *gin.Context) (TableModel, error) {

	var table TableModel

	err := Config.DB.Preload("CategoryContent").Where("ID = ? AND enabled = ?", id, true).Find(&table).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return table, err
}
