package User

import (
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Id         uint      `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Type       string    `json:"type"`
	Image      string    `json:"image"`
	Enabled    bool      `json:"enabled"`
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
		Id:    ss.user.Id,
		Username:  ss.user.Username,
		Email:  ss.user.Email,
		Password:  ss.user.Password,
		Type:  ss.user.Type,
		Image:  ss.user.Image,
		Enabled:  ss.user.Enabled,
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
