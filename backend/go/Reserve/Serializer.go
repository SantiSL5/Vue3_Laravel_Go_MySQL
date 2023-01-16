package Reserve

import (
	User   "namazu/User"
	Table  "namazu/Table"

	"github.com/gin-gonic/gin"
)

type ReserveResponse struct {
	Id              uint                   `json:"id"`
	User            uint                   `json:"user"`
	Table           uint                   `json:"table"`
	Date            string                 `json:"date"`
	Turn            string                 `json:"turn"`
	Confirmed       bool                   `json:"confirmed"`
	UserContent     User.UserModel         `json:"usercontent"`
	TableContent    Table.TableModel       `json:"tablecontent"`
}

type ReserveSerializer struct {
	C     *gin.Context
	reserve ReserveModel
}

type ReservesSerializer struct {
	C      *gin.Context
	reserves []ReserveModel
}

func (ss *ReserveSerializer) Response() ReserveResponse {
	response := ReserveResponse{
		Id:              ss.reserve.Id,
		User:            ss.reserve.User,
		Table:           ss.reserve.Table,
		Date:            ss.reserve.Date,
		Turn:            ss.reserve.Turn,
		Confirmed:       ss.reserve.Confirmed,
		UserContent:     ss.reserve.UserContent,
		TableContent:    ss.reserve.TableContent,
	}

	return response
}

func (ss *ReservesSerializer) Response() []ReserveResponse {
	response := []ReserveResponse{}

	for _, reserve := range ss.reserves {
		serializer := ReserveSerializer{ss.C, reserve}
		response = append(response, serializer.Response())
	}

	return response
}
