package User

import (
	// "fmt"
	"namazu/Config"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRepo(user *UserModel, c *gin.Context) (err error, exist bool) {
	err = Config.DB.Create(user).Error
	return err, false
}

// func UserLoginRepo(usrinput *UserModel, c *gin.Context) (err error, user UserModel) {
	
// 	// err = Config.DB.Where("ID = ? AND enabled = ?", id, true).Scan(&user).Error
// 	// // err = Config.DB.Raw("SELECT * FROM users u WHERE u.email = ? AND u.password = ?", usrinput.Email, usrinput.Password).Scan(&user).Error
// 	// fmt.Println(user)
// 	// return err, user
// }

func CheckUserEmail(user *UserModel, c *gin.Context) (exists UserModel, err error) {
	err = Config.DB.Where("email = ?", user.Email).Find(&exists).Error
	return exists, err
}