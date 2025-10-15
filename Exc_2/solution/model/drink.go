package model

type Drink struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	// todo Add fields: Name, Price, Description with suitable types
	// todo json attributes need to be snakecase, i.e. name, created_at, my_variable, ..
}
