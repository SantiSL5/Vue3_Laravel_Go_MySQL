package DishType

type DishTypeModel struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Order int    `json:"order"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}

func (b *DishTypeModel) TableName() string {
	return "dish_types"
}
