package User

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterService(c *gin.Context) (error, bool) {
	var usrModel UserModel
	c.BindJSON(&usrModel)
	usrModel.Type = "client"
	usrModel.Enabled = true
	usrModel.Image = "https://static.productionready.io/images/smiley-cyrus.jpg"
	usrModel.setPassword(usrModel.Password)
	exists, err := CheckUserEmail(&usrModel, c)

	if exists.Id != 0 {
		println("exist")
		return err, true
	} else {
		return RegisterRepo(&usrModel, c)
	}
}

func LoginService(c *gin.Context) (error, UserModel) {
	var usrModel UserModel
	c.BindJSON(&usrModel)
	exists, err := CheckUserEmail(&usrModel, c)
	if exists.Id != 0 {
		fmt.Println(exists)
		if exists.checkPassword(usrModel.Password) != nil {
			usrModel.clean()
			return err, usrModel
		} else {
			return err, exists
		}
	}
	return err, exists
	// usrModel.setPassword(usrModel.Password)
}

func GetUserServiceByID(c *gin.Context, id uint) (UserModel, error) {
	return GetOneUserByID(id, c)
}
