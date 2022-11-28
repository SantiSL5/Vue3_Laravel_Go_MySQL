package Category

import (
	"github.com/jinzhu/gorm"
	"namazu/Config"
)

type Category struct {
	gorm.Model
	Id      	uint   	`json:"id"`
	Code    	int 	`json:"code"`
	Category   	string 	`json:"category"`
	Capacity   	string 	`json:"capacity"`
	Reserved 	bool 	`json:"reserved"`
}

func (b *Category) CategoryName() string {
	return "category"
}


// //GetAllCategories Fetch all category data
// func GetAllCategories(able *[]Category) (err error) {
// 	if err = Config.DB.Find(category).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

//CreateCategory ... Insert New data
func CreateCategoryModel(category *Category) (err error) {
	if err = Config.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

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
