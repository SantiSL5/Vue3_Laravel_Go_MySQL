package Category

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCategoriesService(c *gin.Context) []CategoryModel {
	return GetAllCategoriesRepo(c)
}

func GetOneCategoryService(id string, c *gin.Context) (CategoryModel, error) {
	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}

	return GetOneCategoryRepo(s, c)
}
