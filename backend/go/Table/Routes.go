package Table

import (
	"github.com/gin-gonic/gin"
)

func TableRouting(router *gin.RouterGroup) {
	router.GET("/", GetTables)
	router.GET(":id", GetTableByID)
}