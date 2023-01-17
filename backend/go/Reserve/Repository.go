package Reserve

import (
	"fmt"
	"namazu/Config"
	"namazu/User"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllReserves Fetch all reserve data
// func GetAllReservesRepo(c *gin.Context) []ReserveModel {

// 	var reserves []ReserveModel

// 	if err := Config.DB.Preload("TableContent").Preload("UserContent").Find(&reserves).Error; err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 		fmt.Println("Status:", err)
// 	}

// 	return reserves
// }

func GetReservesUserRepo(c *gin.Context, u User.UserModel) []ReserveModel {

	var reserves []ReserveModel

	if err := Config.DB.Preload("TableContent").Preload("UserContent").Where("user = ?", u.Id).Find(&reserves).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}
	return reserves
}


func GetOneReserveRepo(id int, c *gin.Context) (ReserveModel, error) {

	var reserve ReserveModel

	err := Config.DB.Preload("TableContent").Preload("UserContent").Where("ID = ?", id).Find(&reserve).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return reserve, err
}

// func CheckReserveRepo(reserve *ReserveModel, c *gin.Context) (ReserveModel, error) {
// 	fmt.Println(reserve.Date)
// 	err := Config.DB.Preload("UserContent").Where("table = ?",reserve.Table).Find(&reserve).Error

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}

// 	return reserve, err
// }

// func CreateReserveRepo(reserve *ReserveModel, c *gin.Context) (err error, exist bool) {
// 	err = Config.DB.Create(reserve).Error
// 	return err, false
// }

func CreateReserveRepo(reserve *ReserveModel, c *gin.Context) (error,bool)  {
	err := Config.DB.Create(reserve).Error
	fmt.Println(err.Error())
	if err != nil {
		fmt.Println(err.Error())
		return err,false
	}

	return err, true
	
}