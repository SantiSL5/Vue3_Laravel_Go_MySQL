package Category

import (
	"github.com/gin-gonic/gin"
)

func GetAllCategoriesService(c *gin.Context) []CategoryModel {
	return GetAllCategoriesDB(c)
}

func GetOneCategoryService(id int, c *gin.Context) (CategoryModel, error) {
	return GetOneCategoryDB(id, c)
}
