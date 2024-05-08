package view

import (
	"fmt"
	"main/model/food"
	"main/model/order"
	"main/utils"
	"strconv"
	"strings"
)

var currentOrder order.Order
var totalPaid int

func SetupOrder(f *food.Food) {
  currentOrder = order.Order{}
	utils.ClearScreen()
	ShowFoodTable()
  fmt.Println("Cancel order [109]\n")

	var qty int  
	fmt.Printf("How many " + strings.Title(f.Name) + " do you want ? ")
	fmt.Scanln(&qty)

  if qty > f.Stock{
    utils.SendMessage("Not enough stock !", ShowFoodMenu)
    return
  }

  if qty == 109 {
    ShowFoodMenu()
    return
  }

  currentOrder = order.Order{Food: f, Quantity: qty}

	fmt.Println(
  	"\n" +
    strings.Title(f.Name) +
    " x" +
    strconv.Itoa(qty) +
    utils.CreateLine(4, " ") +
    utils.ToCurrency(qty*f.Price*1000) +
    "\n")

  fmt.Printf("Are you sure [y/n] ? ")
  utils.ReceiveUserInput(handleOrderInput)
}

func handleOrderInput(input string){
  input = strings.ToLower(input)
  switch input {
    case "y", "yes":
      onOrderConfirmed()
    case "n", "no":
      ShowFoodMenu()
    default:
      utils.SendMessage(nil, ShowFoodMenu)
  }  
}

func ShowOrderDetails(){
  orderTable := ShowOrderTable("Orders")
  fmt.Println(orderTable)
  fmt.Printf(utils.CreateLine(50,"="))

  totalText := "Total"
  totalStr := utils.ToCurrency(totalPaid)
                           
  fmt.Println()
  fmt.Println(totalText + utils.CreateLine(50 - len(totalText+totalStr)," ") + totalStr)
  
  fmt.Printf(utils.CreateLine(50,"="))
}

func ShowOrderMenu(){
  utils.ClearScreen()
  ShowOrderDetails()
                           
  menu := "\n\n[1] Input new order"
  menu += "\n[2] Remove order"
  menu += "\n[3] Checkout"
  fmt.Println(menu)
  fmt.Printf("\nEnter the menu you choose : ")
  utils.ReceiveUserInput(handleOrderMenu)
}

func handleOrderMenu(input string){
  intInput, err := strconv.Atoi(input)
  if err != nil {
    utils.SendMessage(nil, ShowOrderMenu)
  }
  switch intInput {
    case 1:
      ShowFoodMenu()
    case 2:
      removeOrderMenu()
    case 3:
      ShowPaymentMenu(totalPaid)
    default:
      utils.SendMessage(nil, ShowOrderMenu)
  }
}

func removeOrderMenu(){
  var orderIndex int
  fmt.Printf("Enter the order number you want to delete : ")
  fmt.Scanln(&orderIndex)

  if orderIndex-1 >= len(order.GetOrders()) || orderIndex-1 < 0{
    utils.SendMessage("Order not found !", ShowOrderMenu)
    return
  }

  orderSelected := order.GetOrders()[orderIndex-1]
  food.UpdateStock(orderSelected.Food, -orderSelected.Quantity)
  
  order.RemoveOrder(orderIndex-1)

  if len(order.GetOrders()) > 0 {
    ShowOrderMenu()
    return
  }

  ShowFoodMenu()
}

func ShowOrderTable(title string) string {
  var numberStr, foodNameList, foodQuantityList, foodPriceList []string
  orders := order.GetOrders()
  totalPaid = 0

  for number, order := range orders{
    orderPrice := order.Quantity * order.Food.Price * 1000
    totalPaid += orderPrice

    numberStr = append(numberStr, strconv.Itoa(number+1))
    foodNameList = append(foodNameList, strings.Title(order.Food.Name))
    foodQuantityList = append(foodQuantityList, "x"+strconv.Itoa(order.Quantity))
    foodPriceList = append(foodPriceList, utils.ToCurrency(orderPrice))
  }

  table := utils.CreateTable(title, 50, 5,numberStr, foodNameList, foodQuantityList, foodPriceList)
  return table
}

func onOrderConfirmed(){
	order.Add(currentOrder)
  food.UpdateStock(currentOrder.Food, currentOrder.Quantity)
  ShowOrderMenu()
}
