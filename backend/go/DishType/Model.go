package DishType

type DishTypeModel struct {
	Id    uint   `json:"id"`
	Type  string `json:"type"`
	Photo string `json:"photo"`
	// Created_at string `json:"created_at"`
	// Updated_at string `json:"updated_at"`
}

func (b *DishTypeModel) TableName() string {
	return "dish_types"
}
