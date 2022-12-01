package Category

import (
	"github.com/gin-gonic/gin"
)

func CategoryRouting(router *gin.RouterGroup) {
	router.GET("/", GetCategories)
	router.GET("/:id", GetCategoryByID)
}