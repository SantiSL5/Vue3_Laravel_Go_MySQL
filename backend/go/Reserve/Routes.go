package Reserve

import (
	"namazu/User"
	"github.com/gin-gonic/gin"
)

func ReserveRouting(router *gin.RouterGroup) {
	router.Use(User.AuthMiddleware(true))
	router.GET("/", GetReserves)
	router.GET(":id", GetReserveByID)
	router.POST("/", CreateReserve)
}
