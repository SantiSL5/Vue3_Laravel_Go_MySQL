package User

import (
	"github.com/gin-gonic/gin"
)

func UserRouting(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.POST("/register", Register)
}
