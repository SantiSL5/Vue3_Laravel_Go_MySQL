package Category

type CategoryModel struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}

func (b *CategoryModel) TableName() string {
	return "categories"
}
