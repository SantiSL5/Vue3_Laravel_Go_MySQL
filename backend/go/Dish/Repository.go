package Dish

import (
	"fmt"
	"namazu/Config"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

//GetAllDishes Fetch all dish data
func GetAllDishesRepo(c *gin.Context) []DishModel {

	var dishes []DishModel

	if err := Config.DB.Preload("TypeContent").Find(&dishes).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	sort.Slice(dishes, func(i, j int) bool {
		return dishes[i].TypeContent.Order < dishes[j].TypeContent.Order
	})

	return dishes
}

func GetOneDishRepo(id int, c *gin.Context) (DishModel, error) {

	var dish DishModel

	err := Config.DB.Preload("TypeContent").Where("ID = ?", id).Find(&dish).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return dish, err
}
