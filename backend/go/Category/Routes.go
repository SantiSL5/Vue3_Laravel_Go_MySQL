package Category

import (
	"github.com/gin-gonic/gin"
)

func CategoryRouting(router *gin.RouterGroup) {
	router.GET("/", GetAllCategories)
	router.GET("/:id", GetCategoryByID)
}
