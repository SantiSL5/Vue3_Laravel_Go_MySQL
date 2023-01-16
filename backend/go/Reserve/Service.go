package Reserve

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllReservesService(c *gin.Context) []ReserveModel {
	return GetAllReservesRepo(c)
}

func GetOneReserveService(id string, c *gin.Context) (ReserveModel, error) {
	s, err := strconv.Atoi(id)
	if err != nil {
		println("error")
	}

	return GetOneReserveRepo(s, c)
}
