package DishType

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetDishTypes ... Get all dishTypes
func GetAllDishTypes(c *gin.Context) {
	var dishTypes []DishTypeModel=GetAllDishTypesService(c)
	serializer := DishTypesSerializer{c, dishTypes}

	c.JSON(http.StatusOK, serializer.Response())
}

func GetDishTypeByID(c *gin.Context) {
	var dishType, err = GetOneDishTypeService(c.Param("id"), c)

	if err != nil {
		c.JSON(http.StatusNotFound, "DishType doesn't exist")
	} else {
		serializer := DishTypeSerializer{c, dishType}
		c.JSON(http.StatusOK, serializer.Response())
	}

}
