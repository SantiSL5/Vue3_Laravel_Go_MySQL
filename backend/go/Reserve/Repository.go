package Reserve

import (
	"fmt"
	"namazu/Config"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllReserves Fetch all reserve data
func GetAllReservesRepo(c *gin.Context) []ReserveModel {

	var reserves []ReserveModel

	if err := Config.DB.Preload("TableContent").Preload("UserContent").Find(&reserves).Error; err != nil {
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
