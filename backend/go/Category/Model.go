package Models

type Category struct {
	Id      	uint   	`json:"id"`
	Code    	int 	`json:"code"`
	Category   	string 	`json:"category"`
	Capacity   	string 	`json:"capacity"`
	Reserved 	bool 	`json:"reserved"`
}

func (b *Table) TableName() string {
	return "table"
}
