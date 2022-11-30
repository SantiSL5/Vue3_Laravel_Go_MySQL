package Category

import (
	"fmt"
	"namazu/Config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryModel struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}

func (b *CategoryModel) TableName() string {
	return "categories"
}

//GetAllCategories Fetch all category data
func GetAllCategoriesDB(c *gin.Context) []CategoryModel {

	var categories []CategoryModel

	if err := Config.DB.Find(&categories).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}

	return categories

}

func GetOneCategoryDB(id int, c *gin.Context) CategoryModel {

	var category CategoryModel

	if err := Config.DB.Where("ID = ?", id).Find(&category).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return category
	} else {
		return category
	}

}

//CreateCategory ... Insert New data
// func CreateCategoryModel(category *Category) (err error) {
// 	if err = Config.DB.Create(category).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// //GetCategoryByID ... Fetch only one category by Id
// func GetCategoryByID(category *Category, id string) (err error) {
// 	if err = Config.DB.Where("id = ?", id).First(category).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// //UpdateCategory ... Update category
// func UpdateCategory(category *Category, id string) (err error) {
// 	fmt.Println(category)
// 	Config.DB.Save(category)
// 	return nil
// }

// //DeleteCategory ... Delete category
// func DeleteCategory(category *Category, id string) (err error) {
// 	Config.DB.Where("id = ?", id).Delete(category)
// 	return nil
// }
