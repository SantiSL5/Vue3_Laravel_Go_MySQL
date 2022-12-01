package Category

import (

	// "namazu/Category/CategoryModel"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetCategories ... Get all categories
func GetAllCategories(c *gin.Context) {
	var categories []CategoryModel

	categories = GetAllCategoriesDB(c)

	println(categories)

	serializer := CategoriesSerializer{c, categories}

	c.JSON(http.StatusOK, serializer.Response())

	// println(categories)
	// err := GetAllCategories(&categories)
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// } else {
	// 	c.JSON(http.StatusOK, category)
	// }
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

//CreateCategory ... Create Category
// func CreateCategory(c *gin.Context) {
// 	var category Category
// 	c.BindJSON(&category)
// 	fmt.Println("Hello World!")

// 	err := CreateCategoryModel(&category)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, category)
// 	}
// }

// //GetCategoryByID ... Get the category by id
// func GetCategoryByID(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var category Category
// 	err := GetCategoryByID(&category, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, category)
// 	}
// }

// //UpdateCategory ... Update the category information
// func UpdateCategory(c *gin.Context) {
// 	var category Category
// 	id := c.Params.ByName("id")
// 	err := GetCategoryByID(&category, id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, category)
// 	}
// 	c.BindJSON(&category)
// 	err = UpdateCategory(&category, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, category)
// 	}
// }

// //DeleteCategory ... Delete the category
// func DeleteCategory(c *gin.Context) {
// 	var category Category
// 	id := c.Params.ByName("id")
// 	err := DeleteCategory(&category, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
// 	}
// }
