package Category

import (
	"fmt"
	"net/http"
	// "namazu/Category"

	"github.com/gin-gonic/gin"
)

// //GetCategories ... Get all categories
// func GetCategories(c *gin.Context) {
// 	var category [] Category
// 	err := GetAllCategories(&category)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, category)
// 	}
// }

//CreateCategory ... Create Category
func CreateCategory(c *gin.Context) {
	var category Category
	c.BindJSON(&category)
	fmt.Println("Hello World!")

	err := CreateCategoryModel(&category)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}
}

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
