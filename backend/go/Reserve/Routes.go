package Reserve

import (
	"namazu/User"
	"github.com/gin-gonic/gin"
)

func ReserveRouting(router *gin.RouterGroup) {
	router.GET("/tables", GetTables)
	router.Use(User.AuthMiddleware(true))
	router.GET("/", GetReservesUser)
	router.GET(":id", GetReserveByID)
	router.POST("/", CreateReserve)
	router.DELETE(":id", CancelReserve)
	router.Use(User.AuthMiddleware(false))
}
