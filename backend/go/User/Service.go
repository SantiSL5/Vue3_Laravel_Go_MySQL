package User

import (
	"fmt"
	"errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func RegisterService(c *gin.Context) (error, bool) {
	var newUser UserModel
	c.BindJSON(&newUser)
	newUser.Type = "client"
	newUser.Enabled = true
	newUser.Image = "https://static.productionready.io/images/smiley-cyrus.jpg"
	if len(newUser.Password) == 0 {
		err:=errors.New("Password required")
		return err,false
	}
	bytePassword := []byte(newUser.Password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	newUser.Password = string(passwordHash)
	exists := CheckUserEmail(&newUser, c)

	if exists.Id != 0 {
		err:=errors.New("Email is already taken")
		return err, true
	} else {
		return RegisterRepo(&newUser, c)
	}
}

func LoginService(c *gin.Context) (error, UserModel) {
	var newUser UserModel
	c.BindJSON(&newUser)
	exists := CheckUserEmail(&newUser, c)
	if exists.Id != 0 {
		fmt.Println(exists)
		bytePassword := []byte(newUser.Password)
		byteHashedPassword := []byte(exists.Password)
		fmt.Println(bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword))
		err:=errors.New("Email is already taken")
		if bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword) != nil {
			// usrModel.clean()
			return err, newUser
		} else {
			return err, exists
		}
	}
	err:=errors.New("Email is not registered")
	return err, exists
	// usrModel.setPassword(usrModel.Password)
}

func GetUserServiceByID(c *gin.Context, id uint) (UserModel, error) {
	return GetOneUserByID(id, c)
}
