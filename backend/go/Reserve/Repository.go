package Reserve

import (
	"fmt"
	"namazu/Config"
	"namazu/User"
	"namazu/Table"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReservesDayTurn(c *gin.Context, date string, turn string) []uint {
	var reserves []ReserveModel
	var ids []uint
	fmt.Println(date)
	if err := Config.DB.Preload("TableContent").Preload("TableContent").Raw("SELECT * FROM reserves WHERE date = ? AND turn = ?", date, turn).Find(&reserves).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}
	for _, reserve := range reserves {
		ids = append(ids, reserve.Table)
	}
	return ids
}

func GetAllTablesRepo(c *gin.Context) (tables []Table.TableModel, count int) {
	var countTables []Table.TableModel
	var ids []uint
	limitV, limit := c.GetQuery("limit")
	offsetV, offset := c.GetQuery("offset")
	capacityV, capacity := c.GetQuery("capacity")
	categoryV, category := c.GetQuery("category")
	dateV, date := c.GetQuery("date")
	turnV, turn := c.GetQuery("turn")
	if date && turn {
		ids=GetReservesDayTurn(c,dateV,turnV)
		if len(ids) == 0 {
			ids=[]uint{0}
		}
		if limit && limitV != "undefined" && offset && capacity {
			count := 0
			if category {
				if err := Config.DB.Preload("CategoryContent").Where("enabled = ? AND capacity >= ? AND category = ? AND id NOT IN(?)", true,  capacityV, categoryV, ids).Limit(limitV).Offset(offsetV).Find(&tables).Error; err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					fmt.Println("Status:", err)
				}
				if err := Config.DB.Preload("CategoryContent").Where("enabled = ? AND capacity >= ? AND category = ? AND id NOT IN(?)", true,  capacityV, categoryV, ids).Find(&countTables).Error; err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					fmt.Println("Status:", err)
				}
			}else {
				if err := Config.DB.Preload("CategoryContent").Where("enabled = ? AND capacity >= ? AND id NOT IN(?)", true,  capacityV, ids).Limit(limitV).Offset(offsetV).Find(&tables).Error; err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					fmt.Println("Status:", err)
				}
				if err := Config.DB.Preload("CategoryContent").Where("enabled = ? AND capacity >= ? AND id NOT IN(?)", true, capacityV, ids).Find(&countTables).Error; err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					fmt.Println("Status:", err)
				}
			}
			count= len(countTables)
			return tables, count
		} else {
			if err := Config.DB.Preload("CategoryContent").Where("enabled = ?", true).Find(&tables).Error; err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				fmt.Println("Status:", err)
			}
			if err := Config.DB.Preload("CategoryContent").Where("enabled = ?", true).Find(&countTables).Error; err != nil {
				c.AbortWithStatus(http.StatusNotFound)
				fmt.Println("Status:", err)
			}
			count= len(countTables)
			return tables, count
		}

	}
	return tables, count
}


func GetReservesUserRepo(c *gin.Context, u User.UserModel) []ReserveModel {

	var reserves []ReserveModel

	if err := Config.DB.Preload("TableContent").Preload("UserContent").Where("user = ?", u.Id).Find(&reserves).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println("Status:", err)
	}
	return reserves
}


func GetOneReserveRepo(id int, c *gin.Context) (ReserveModel, error) {

	var reserve ReserveModel

	err := Config.DB.Preload("TableContent").Preload("UserContent").Where("ID = ?", id).Find(&reserve).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	return reserve, err
}

func CreateReserveRepo(reserve *ReserveModel, c *gin.Context) (error,bool)  {
	err := Config.DB.Create(reserve).Error
	fmt.Println(err.Error())
	if err != nil {
		fmt.Println(err.Error())
		return err,false
	}

	return err, true
	
}

func CancelReserveRepo(id int, c *gin.Context) error  {
	var reserve ReserveModel
	Config.DB.First(&reserve, id)
	err := Config.DB.Delete(&reserve).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}