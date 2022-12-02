package Table

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTables(c *gin.Context) {
	var tables []TableModel

	tables = GetAllTablesDB(c)

	serializer := TablesSerializer{c, tables}

	c.JSON(http.StatusOK, serializer.Response())
}

func GetTableByID(c *gin.Context) {
	s := c.Param("id")
	var table TableModel
	var id int
	id, err := strconv.Atoi(s)
	if err != nil {
		println("error")
	}

	table, err = GetOneTableDB(id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, "Table doesn't exist")
	} else {
		serializer := TableSerializer{c, table}
		c.JSON(http.StatusOK, serializer.Response())
	}
}
