package Dish

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDishesService(c *gin.Context) []DishModel {
	return GetAllDishesRepo(c)
}

func GetOneDishService(id string, c *gin.Context) (DishModel, error) {
	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}

	return GetOneDishRepo(s, c)
}
