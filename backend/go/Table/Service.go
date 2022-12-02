package Table

import (
	"github.com/gin-gonic/gin"
)

func GetAllTablesService(c *gin.Context) []TableModel {
	return GetAllTablesDB(c)
}

func GetOneTableService(id int, c *gin.Context) (TableModel, error) {
	return GetOneTableDB(id, c)
}
