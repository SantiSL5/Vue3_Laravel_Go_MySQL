package Dish

import (
	DishType "namazu/DishType"

	"github.com/gin-gonic/gin"
)

type DishResponse struct {
	Id              uint                   `json:"id"`
	Dish            string                 `json:"dish"`
	Type            uint                   `json:"type"`
	Price           float64                `json:"price"`
	Desc            string                 `json:"desc"`
	Photo        	string                 `json:"photo"`
	TypeContent     DishType.DishTypeModel `json:"typecontent"`
}

type DishSerializer struct {
	C     *gin.Context
	dish  DishModel
}

type DishesSerializer struct {
	C      *gin.Context
	dishes []DishModel
}

func (ss *DishSerializer) Response() DishResponse {
	response := DishResponse{
		Id:              ss.dish.Id,
		Dish:            ss.dish.Dish,
		Type:            ss.dish.Type,
		Price:           ss.dish.Price,
		Desc:            ss.dish.Desc,
		Photo:           ss.dish.Photo,
		TypeContent:     ss.dish.TypeContent,
	}

	return response
}

func (ss *DishesSerializer) Response() []DishResponse {
	response := []DishResponse{}

	for _, dish := range ss.dishes {
		serializer := DishSerializer{ss.C, dish}
		response = append(response, serializer.Response())
	}

	return response
}
