package Table

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOneTableService(id string, c *gin.Context) (TableModel, error) {
	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}

	return GetOneTableRepo(s, c)
}
