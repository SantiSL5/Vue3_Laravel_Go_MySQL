package Category

import (
	"fmt"
	"namazu/Config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategoriesDB(c *gin.Context) []CategoryModel {

	var categories []CategoryModel

	if err := Config.DB.Find(&categories).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return categories

}

func GetOneCategoryDB(id int, c *gin.Context) (CategoryModel, error) {

	var category CategoryModel

	err := Config.DB.Where("ID = ?", id).Find(&category).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return category, err

}
