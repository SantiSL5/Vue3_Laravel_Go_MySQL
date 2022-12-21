package DishType

import (
	"fmt"
	"namazu/Config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDishTypesRepo(c *gin.Context) []DishTypeModel {

	var dishTypes []DishTypeModel
	limitV, limit := c.GetQuery("limit")
	offsetV, offset := c.GetQuery("offset")

	if limit && limitV != "undefined" && offset {
		if err := Config.DB.Limit(limitV).Offset(offsetV).Find(&dishTypes).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println("Status:", err)
		}
	} else {
		if err := Config.DB.Find(&dishTypes).Error; err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			fmt.Println("Status:", err)
		}
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
