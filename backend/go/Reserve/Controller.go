package Reserve

import (
	"net/http"
	"namazu/Table"

	"github.com/gin-gonic/gin"
)

func GetTables(c *gin.Context) {
	var tablas []Table.TableModel
	var count int
	tablas, count = GetAllTablesService(c)
	serializer := Table.TablesSerializer{c, tablas}

	c.JSON(http.StatusOK,gin.H{
	  	"tables":serializer.Response(), 
		"count": count,
	})
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

func GetReservesUser(c *gin.Context) {
	var reserves []ReserveModel = GetReservesUserService(c)
	serializer := ReservesSerializer{c, reserves}

	c.JSON(http.StatusOK, serializer.Response())
}

func CreateReserve(c *gin.Context) {
	err,result:=CreateReserveService(c)
	if err != nil || result {
		c.JSON(http.StatusOK, "Reserve created correctly")	
	}
}