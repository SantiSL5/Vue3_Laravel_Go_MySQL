package User

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	// var category, err = GetOneCategoryService(c.Param("id"), c)

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, "Category doesn't exist")
	// } else {
	// 	serializer := CategorySerializer{c, category}
	// 	c.JSON(http.StatusOK, serializer.Response())
	// }

}

func Register(c *gin.Context) {
	err, bool := RegisterService(c)
	if err != nil || bool {
		c.JSON(http.StatusInternalServerError, "Email is registered")
	} else {
		c.JSON(http.StatusOK, "User created correctly")
	}
}

func Login(c *gin.Context) {
	// var category, err = GetOneCategoryService(c.Param("id"), c)

	// if err != nil {
	// 	c.JSON(http.StatusNotFound, "Category doesn't exist")
	// } else {
	// 	serializer := CategorySerializer{c, category}
	// 	c.JSON(http.StatusOK, serializer.Response())
	// }

}
