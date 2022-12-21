package Dish

import (
	"github.com/gin-gonic/gin"
)

func DishRouting(router *gin.RouterGroup) {
	router.GET("/", GetDishes)
	router.GET(":id", GetDishByID)
}
