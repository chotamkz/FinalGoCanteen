package order

import (
	"main/model/food"
	"encoding/json"
    "io/ioutil"
    "log"
    "os"
)

type OrderHistory struct {
	Orders []Order `json:"orders"`
  }

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

////

  
var orderHistory OrderHistory
  
func AddOrder(order Order) {
	orderHistory.Orders = append(orderHistory.Orders, order)
	saveOrderHistory()
}
  
func Init() {
	data, err := ioutil.ReadFile("DB/orderHistory.json")
	if err != nil {
		log.Fatal(err)
	}
  
	json.Unmarshal(data, &orderHistory)
}
  
func saveOrderHistory() {
	json, err := json.Marshal(orderHistory)
	if err != nil {
		log.Fatal(err)
	}
  
	os.WriteFile("DB/orderHistory.json", json, 0666)
}