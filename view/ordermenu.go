package view

import (
	"fmt"
	"main/model/food"
	"main/model/order"
  "main/model/orderHistory"
	"main/utils"
	"strconv"
	"strings"
  "log"
  "os"
  "encoding/json"
  "encoding/csv"
)

var currentOrder order.Order
var totalPaid int

func SetupOrder(f *food.Food) {
  currentOrder = order.Order{}
	utils.ClearScreen()
	ShowFoodTable()
  fmt.Println("Cancel order [109]\n")

	var qty int  
	fmt.Printf("How many " + strings.Title(f.Name) + " do you want? ")
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
    utils.ToCurrency(qty*f.Price) +
    "\n")

  fmt.Printf("Are you sure [y/n] ? ")
  utils.ReceiveUserInput(handleOrderInput)
}

func handleOrderInput(input string){
  input = strings.ToLower(input)
  switch input {
    case "y", "yes", "da":
      onOrderConfirmed()
    case "n", "no", "net":
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
  menu += "\n[3] Proceed to payment"
  fmt.Println(menu)
  fmt.Printf("\nEnter the menu you choose: ")
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
    orderPrice := order.Quantity * order.Food.Price
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
  addToOrderHistory(currentOrder)
  ShowOrderMenu()
}


func addToOrderHistory(order order.Order) {
  file, err := os.ReadFile("DB/orderHistory.json")
  if err != nil {
      log.Fatal(err)
  }

  var orderHistory orderHistory.OrderHistory
  if err := json.Unmarshal(file, &orderHistory); err != nil {
      log.Fatal(err)
  }

  orderHistory.Orders = append(orderHistory.Orders, order)

  // updatedJSON, err := json.Marshal(orderHistory)
  // if err != nil {
  //     log.Fatal(err)
  // }

  // if err := os.WriteFile("DB/orderHistory.json", updatedJSON, 0644); err != nil {
  //     log.Fatal(err)
  // }

  csvFile, err := os.OpenFile("DB/orderHistory.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
      log.Fatal(err)
  }
    defer csvFile.Close()

    // Create a CSV writer
    writer := csv.NewWriter(csvFile)
    defer writer.Flush()

    fileInfo, err := csvFile.Stat()
    if err != nil {
        log.Fatal(err)
    }
    if fileInfo.Size() == 0 {
        if err := writer.Write([]string{"Name", "Price", "Stock"}); err != nil {
            log.Fatal(err)
        }
    }

    if err := writer.Write([]string{
        order.Food.Name,
        strconv.Itoa(order.Food.Price),
        strconv.Itoa(order.Quantity),
    }); err != nil {
        log.Fatal(err)
    }
}