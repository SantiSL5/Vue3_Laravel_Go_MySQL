package User

import (
	"errors"
	// "fmt"
	// "os"
	// "time"

	// "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Id         uint      `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Type       string    `json:"type"`
	Image      string    `json:"image"`
	Enabled    bool      `json:"enabled"`
}

func (b *UserModel) TableName() string {
	return "users"
}

// func (u *UserModel) clean() error {
// 	u.Id = 0
// 	u.Email = ""
// 	u.Name = ""
// 	u.Img = ""
// 	u.Enabled = false
// 	u.Password = ""
// 	return nil
// }

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password required")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.Password = string(passwordHash)
	return nil
}

// func (u *UserModel) checkPassword(password string) error {
// 	bytePassword := []byte(password)
// 	byteHashedPassword := []byte(u.Password)
// 	fmt.Println(bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword))
// 	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
// }

// func GenToken(id uint) string {
// 	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
// 	jwt_token.Claims = jwt.MapClaims{
// 		"id":  id,
// 		"exp": time.Now().Add(time.Hour * 24).Unix(),
// 	}
// 	token, _ := jwt_token.SignedString([]byte(os.Getenv("SecretPassword")))
// 	return token
// }