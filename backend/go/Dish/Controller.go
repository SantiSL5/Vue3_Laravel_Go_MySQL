package Dish

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDishes(c *gin.Context) {
	var dishes []DishModel = GetAllDishesService(c)
	serializer := DishesSerializer{c, dishes}

	c.JSON(http.StatusOK, serializer.Response())
}

func GetDishByID(c *gin.Context) {
	var dish, err = GetOneDishService(c.Param("id"), c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Dish doesn't exist")
	} else {
		serializer := DishSerializer{c, dish}
		c.JSON(http.StatusOK, serializer.Response())
	}
}
