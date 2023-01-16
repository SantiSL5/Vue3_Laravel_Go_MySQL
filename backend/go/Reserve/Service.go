package Reserve

import (
	"strconv"
	"namazu/User"
	// "namazu/Table"
	// "fmt"

	"github.com/gin-gonic/gin"
)

func GetAllReservesService(c *gin.Context) []ReserveModel {
	return GetAllReservesRepo(c)
}

func GetOneReserveService(id string, c *gin.Context) (ReserveModel, error) {
	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}

	return GetOneReserveRepo(s, c)
}

func CreateReserveService(c *gin.Context) (error, bool) {
	var newReserve ReserveModel
	// var newUser User.UserModel
	c.BindJSON(&newReserve)
	usr, _ := c.Get("user_model")
	u, valid := usr.(User.UserModel)
	if valid {
		newReserve.User= u.Id
		newReserve.Confirmed= false	
	}

	err,reserve := CreateReserveRepo(&newReserve, c)

	return err, reserve
}
