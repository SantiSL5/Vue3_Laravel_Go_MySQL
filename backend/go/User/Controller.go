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
		c.JSON(http.StatusInternalServerError, "Email or password is not correct")
	} else {
		UpdateContextUserModel(c, usr.Id)
		// c.Set("my_user_id", user.Id)
		// c.Set("my_user_model", user)
		serializer := UserSerializer{c, usr}
		c.JSON(http.StatusOK, serializer.Response())
		// c.JSON(http.StatusOK, usr)
	}
}
