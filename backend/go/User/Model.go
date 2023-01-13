package User

import (

	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

)

type UserModel struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Image    string `json:"image"`
	Enabled  bool   `json:"enabled"`
}

func (b *UserModel) TableName() string {
	return "users"
}

func GenToken(id uint) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token, _ := jwt_token.SignedString([]byte(os.Getenv("SECRET")))
	return token
}
