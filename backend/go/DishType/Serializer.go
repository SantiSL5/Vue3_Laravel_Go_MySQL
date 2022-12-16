package DishType

import (
	"github.com/gin-gonic/gin"
)

type DishTypeResponse struct {
	Id    uint   `json:"id"`
	Type  string `json:"type"`
	Photo string `json:"photo"`
}

type DishTypeSerializer struct {
	C        *gin.Context
	dishType DishTypeModel
}

type DishTypesSerializer struct {
	C          *gin.Context
	dishTypes []DishTypeModel
}

func (ss *DishTypeSerializer) Response() DishTypeResponse {
	response := DishTypeResponse{
		Id:    ss.dishType.Id,
		Type:  ss.dishType.Type,
		Photo: ss.dishType.Photo,
	}

	return response
}

func (ss *DishTypesSerializer) Response() []DishTypeResponse {
	response := []DishTypeResponse{}

	for _, dishType := range ss.dishTypes {
		serializer := DishTypeSerializer{ss.C, dishType}
		response = append(response, serializer.Response())
	}

	return response
}