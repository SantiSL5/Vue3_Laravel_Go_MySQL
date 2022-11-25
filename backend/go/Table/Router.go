package Table

import (
	"github.com/gin-gonic/gin"
)

func ReservationRouting(router *gin.RouterGroup) {
	router.GET("/:id", GetTable)
	router.GET("/", GetAllTables)
	router.POST("/",CreateTable)
	router.PUT("/:id",UpdateTable)
	router.DELETE("/:id",DeleteTable)
}
