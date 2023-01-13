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
	err, result := RegisterService(c)
	if err != nil || result {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, "User created correctly")
	}
}

func Login(c *gin.Context) {
	err, usr := LoginService(c)
	if err != nil || len(usr.Username) == 0 {
		c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		UpdateContextUserModel(c, usr.Id)
		serializer := UserSerializer{c, usr}
		c.JSON(http.StatusOK, serializer.Response())
	}
}
