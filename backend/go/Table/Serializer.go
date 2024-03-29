package Table

import (
	Category "namazu/Category"

	"github.com/gin-gonic/gin"
)

type TableResponse struct {
	Id              uint                   `json:"id"`
	Category        uint                   `json:"category"`
	Code            int                    `json:"code"`
	Capacity        int                    `json:"capacity"`
	Enabled         bool                   `json:"enabled"`
	CategoryContent Category.CategoryModel `json:"categorycontent"`
}

type TableSerializer struct {
	C     *gin.Context
	table TableModel
}

type TablesSerializer struct {
	C      *gin.Context
	Tables []TableModel
}

func (ss *TableSerializer) Response() TableResponse {
	response := TableResponse{
		Id:              ss.table.Id,
		Code:            ss.table.Code,
		Category:        ss.table.Category,
		Capacity:        ss.table.Capacity,
		Enabled:         ss.table.Enabled,
		CategoryContent: ss.table.CategoryContent,
	}

	return response
}

func (ss *TablesSerializer) Response() []TableResponse {
	response := []TableResponse{}

	for _, table := range ss.Tables {
		serializer := TableSerializer{ss.C, table}
		response = append(response, serializer.Response())
	}

	return response
}
