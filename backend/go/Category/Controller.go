package Category

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetCategories ... Get all categories
func GetAllCategories(c *gin.Context) {
	var categories []CategoryModel

	categories = GetAllCategoriesService(c)

	println(categories)

	serializer := CategoriesSerializer{c, categories}

	c.JSON(http.StatusOK, serializer.Response())
}

func GetCategoryByID(c *gin.Context) {

	s := c.Param("id")
	var category CategoryModel
	var id int
	id, err := strconv.Atoi(s)
	if err != nil {
		println("error")
	}

	category, err = GetOneCategoryDB(id, c)

	if err != nil {
		c.JSON(http.StatusNotFound, "Category doesn't exist")
	} else {
		serializer := CategorySerializer{c, category}
		c.JSON(http.StatusOK, serializer.Response())
	}

}
