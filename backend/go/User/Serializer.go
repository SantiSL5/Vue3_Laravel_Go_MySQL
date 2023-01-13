package User

import (
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Type       string    `json:"type"`
	Image      string    `json:"image"`
	Token      string    `json:"token"`
}

type UserSerializer struct {
	C        *gin.Context
	user UserModel
}

type UsersSerializer struct {
	C          *gin.Context
	users []UserModel
}

func (ss *UserSerializer) Response() UserResponse {
	response := UserResponse{
		Username:  ss.user.Username,
		Email:  ss.user.Email,
		Type:  ss.user.Type,
		Image:  ss.user.Image,
		Token:  GenToken(ss.user.Id),
	}

	return response
}

func (ss *UsersSerializer) Response() []UserResponse {
	response := []UserResponse{}

	for _, user := range ss.users {
		serializer := UserSerializer{ss.C, user}
		response = append(response, serializer.Response())
	}

	return response
}
