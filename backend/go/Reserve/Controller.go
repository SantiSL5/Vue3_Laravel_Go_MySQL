package Reserve

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReserves(c *gin.Context) {
	var reserves []ReserveModel = GetAllReservesService(c)
	serializer := ReservesSerializer{c, reserves}

	c.JSON(http.StatusOK, serializer.Response())
}

func GetReserveByID(c *gin.Context) {
	var reserve, err = GetOneReserveService(c.Param("id"), c)
	if err != nil {
		c.JSON(http.StatusNotFound, "Reserve doesn't exist")
	} else {
		serializer := ReserveSerializer{c, reserve}
		c.JSON(http.StatusOK, serializer.Response())
	}
}

func CreateReserve(c *gin.Context) {
	err, result := CreateReserveService(c)
	if err != nil || result {
		c.JSON(http.StatusOK, "Reserve created correctly")	
	}
}