package order

import (
	"main/model/food"
)

type Order struct {
	Food     *food.Food
	Quantity int
}

var orders []Order

func GetOrders() []Order {
	return orders
}

func Add(o Order) {
	orders = append(orders, o)
}

func RemoveOrder(indexToRemove int) {
	orders = append(orders[:indexToRemove], orders[indexToRemove+1:]...)
}

func ClearOrder() {
	orders = []Order{}
}
