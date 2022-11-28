package Table

import (
	"fmt"
	"net/http"
	// "namazu/Models"

	"github.com/gin-gonic/gin"
)

//GetTables ... Get all tables
func GetTables(c *gin.Context) {
	var table []Models.Table
	err := Models.GetAllTables(&table)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
}

//CreateTable ... Create Table
func CreateTable(c *gin.Context) {
	var table Models.Table
	c.BindJSON(&table)
	err := Models.CreateTable(&table)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
}

//GetTableByID ... Get the table by id
func GetTableByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var table Models.Table
	err := Models.GetTableByID(&table, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
}

//UpdateTable ... Update the table information
func UpdateTable(c *gin.Context) {
	var table Models.Table
	id := c.Params.ByName("id")
	err := Models.GetTableByID(&table, id)
	if err != nil {
		c.JSON(http.StatusNotFound, table)
	}
	c.BindJSON(&table)
	err = Models.UpdateTable(&table, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, table)
	}
}

//DeleteTable ... Delete the table
func DeleteTable(c *gin.Context) {
	var table Models.Table
	id := c.Params.ByName("id")
	err := Models.DeleteTable(&table, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
