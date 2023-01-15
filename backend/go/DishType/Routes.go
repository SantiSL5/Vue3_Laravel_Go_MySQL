package DishType

import (
	"github.com/gin-gonic/gin"
)

func DishTypeRouting(router *gin.RouterGroup) {
	router.GET("/", GetAllDishTypes)
	router.GET(":id", GetDishTypeByID)
}
