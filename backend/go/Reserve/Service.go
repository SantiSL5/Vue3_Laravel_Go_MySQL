package Reserve

import (
	"strconv"
	"errors"
	"fmt"
	"namazu/User"
	"namazu/Table"

	"github.com/gin-gonic/gin"
)

func GetAllTablesService(c *gin.Context) ([]Table.TableModel,int) {
	var tables []Table.TableModel
	var count int
	tables,count=GetAllTablesRepo(c)
	return tables,count
}

func GetReservesUserService(c *gin.Context) []ReserveModel {
	usr, _ := c.Get("user_model")
	u, valid := usr.(User.UserModel)
	if valid {
		return GetReservesUserRepo(c,u)
	}
	return nil
}

func GetOneReserveService(id string, c *gin.Context) (ReserveModel, error) {
	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}

	return GetOneReserveRepo(s, c)
}

func CreateReserveService(c *gin.Context) (error, bool){
	var newReserve ReserveModel
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

func CancelReserveService(id string, c *gin.Context) error {

	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}
	reserve,err:=GetOneReserveRepo(s, c)
	if err != nil {
		fmt.Println(reserve)
		return errors.New("Reserve doesn't exist")
	}
	usr, _ := c.Get("user_model")
	u, valid := usr.(User.UserModel)
	if valid {
		if reserve.User == u.Id {
			err=CancelReserveRepo(s,c)
			return err
		}
		return errors.New("Is not your reserve")
	}

	return err
}
