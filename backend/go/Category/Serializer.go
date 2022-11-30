package Category

import (
	"github.com/gin-gonic/gin"
)

type CategoryResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type CategorySerializer struct {
	C        *gin.Context
	category CategoryModel
}

type CategoriesSerializer struct {
	C          *gin.Context
	categories []CategoryModel
}

func (ss *CategorySerializer) Response() CategoryResponse {
	response := CategoryResponse{
		Id:    ss.category.Id,
		Name:  ss.category.Name,
		Photo: ss.category.Photo,
	}

	return response
}

func (ss *CategoriesSerializer) Response() []CategoryResponse {
	response := []CategoryResponse{}

	for _, category := range ss.categories {
		serializer := CategorySerializer{ss.C, category}
		response = append(response, serializer.Response())
	}

	return response
}
