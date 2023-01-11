package User

import (
	// "strconv"

	"github.com/gin-gonic/gin"
)

func RegisterService(c *gin.Context) (error,bool) {
	var usrModel UserModel
	c.BindJSON(&usrModel)
	usrModel.Type = "client"
	usrModel.Enabled = true
	usrModel.setPassword(usrModel.Password)
	exists, err := CheckUserEmail(&usrModel, c)
	
	if exists.Id != 0 {
		println("exist")
		return err, true
	} else {
		return RegisterRepo(&usrModel, c)
	}
}
