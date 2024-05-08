package orderHistory

import (
	"encoding/json"
	"main/model/order"
    "io/ioutil"
    "log"
    "os"
)

type OrderHistory struct {
	Orders [] order.Order `json:"orders"`
}
  
var orderHistory OrderHistory
  
func AddOrder(order order.Order) {
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