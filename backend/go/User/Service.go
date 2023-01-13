package User

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"

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
	var userlog UserModel
	c.BindJSON(&userlog)
	exists := CheckUserEmail(&userlog, c)
	if exists.Id != 0 {
		bytePassword := []byte(userlog.Password)
		byteHashedPassword := []byte(exists.Password)
		// c.JSON(http.StatusOK, exists.Password)
		result:=bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
		if result != nil {
			err:=errors.New("Incorrect password")
			return err, userlog
		} else {
			err:=result
			return err, exists
		}
		
	}
	err:=errors.New("Email is not registered")
	return err, exists
}

func GetUserServiceByID(c *gin.Context, id uint) (UserModel, error) {
	return GetOneUserByID(id, c)
}
