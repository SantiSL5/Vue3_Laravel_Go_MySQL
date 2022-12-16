package DishType

import (
	"fmt"
	"namazu/Config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDishTypesRepo(c *gin.Context) []DishTypeModel {

	var dishTypes []DishTypeModel

	if err := Config.DB.Find(&dishTypes).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return dishTypes

}

func GetOneDishTypeRepo(id int, c *gin.Context) (DishTypeModel, error) {

	var dishType DishTypeModel

	err := Config.DB.Where("ID = ?", id).Find(&dishType).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return dishType, err

}
