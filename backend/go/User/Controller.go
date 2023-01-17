package User

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	usr, _ := c.Get("user_model")
	u, valid := usr.(UserModel)
	if valid {
		serializer := UserSerializer{c, u}
		c.JSON(http.StatusOK, serializer.Response())
	} else {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	
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
		c.JSON(http.StatusUnauthorized, err.Error())
	} else {
		UpdateContextUserModel(c, usr.Id)
		serializer := UserSerializer{c, usr}
		c.JSON(http.StatusOK, serializer.Response())
	}
}
