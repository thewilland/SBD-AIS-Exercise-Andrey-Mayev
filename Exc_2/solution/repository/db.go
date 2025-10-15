package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	// drinks :=

	drinks := []model.Drink{{ID: 1, Name: "Water", Price: 1.80, Description: "Yay water"},
		{ID: 2, Name: "Evil Water", Price: 6.66, Description: "AAAA OH NO"},
		{ID: 3, Name: "Leaf Water Hot", Price: 2.80, Description: "Oh that's tea mmmm AAAAA HOT OH NO"}}

	orders := []model.Order{{DrinkID: 1, CreatedAt: time.Now(), Amount: 1},
		{DrinkID: 2, CreatedAt: time.Now(), Amount: 13},
		{DrinkID: 3, CreatedAt: time.Now(), Amount: 101}}

	// Init orders slice with some test data

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += order.Amount
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
	//if order == nil {
	//return
	//}
	db.orders = append(db.orders, *order)
}
