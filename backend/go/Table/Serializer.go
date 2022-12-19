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
	Reserved        bool                   `json:"reserved"`
	CategoryContent Category.CategoryModel `json:"categorycontent"`
}

type TableSerializer struct {
	C     *gin.Context
	table TableModel
}

type TablesSerializer struct {
	C      *gin.Context
	tables []TableModel
}

func (ss *TableSerializer) Response() TableResponse {
	response := TableResponse{
		Id:              ss.table.Id,
		Code:            ss.table.Code,
		Category:        ss.table.Category,
		Capacity:        ss.table.Capacity,
		Reserved:        ss.table.Reserved,
		CategoryContent: ss.table.CategoryContent,
	}

	return response
}

func (ss *TablesSerializer) Response() []TableResponse {
	response := []TableResponse{}

	for _, table := range ss.tables {
		serializer := TableSerializer{ss.C, table}
		response = append(response, serializer.Response())
	}

	return response
}
