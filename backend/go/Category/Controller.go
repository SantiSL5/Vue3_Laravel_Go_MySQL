package Category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetCategories ... Get all categories
func GetAllCategories(c *gin.Context) {
	var categories []CategoryModel = GetAllCategoriesService(c)
	serializer := CategoriesSerializer{c, categories}

	c.JSON(http.StatusOK, serializer.Response())
}

func GetCategoryByID(c *gin.Context) {
	var category, err = GetOneCategoryService(c.Param("id"), c)

	if err != nil {
		c.JSON(http.StatusNotFound, "Category doesn't exist")
	} else {
		serializer := CategorySerializer{c, category}
		c.JSON(http.StatusOK, serializer.Response())
	}

}
