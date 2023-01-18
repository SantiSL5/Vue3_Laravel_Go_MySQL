package Table

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTableByID(c *gin.Context) {
	var table, err = GetOneTableService(c.Param("id"), c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Table doesn't exist")
	} else {
		serializer := TableSerializer{c, table}
		c.JSON(http.StatusOK, serializer.Response())
	}
}
