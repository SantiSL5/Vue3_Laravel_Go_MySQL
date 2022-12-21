package DishType

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDishTypesService(c *gin.Context) []DishTypeModel {
	return GetAllDishTypesRepo(c)
}

func GetOneDishTypeService(id string, c *gin.Context) (DishTypeModel, error) {
	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}

	return GetOneDishTypeRepo(s, c)
}
