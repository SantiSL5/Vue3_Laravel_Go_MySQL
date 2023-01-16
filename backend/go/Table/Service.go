package Table

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTablesService(c *gin.Context) ([]TableModel,int) {
	var tables []TableModel
	var count int
	tables,count=GetAllTablesRepo(c)
	return tables,count
}

func GetOneTableService(id string, c *gin.Context) (TableModel, error) {
	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}

	return GetOneTableRepo(s, c)
}
