package view

import (
	"fmt"
	"main/model/food"
	"main/utils"
	"sort"
	"strconv"
	"strings"
)

func ShowFoodMenu() {
	utils.ClearScreen()
	ShowFoodTable()
  fmt.Println("See order details [101]")
  fmt.Println("Back to main menu [109]\n")
	fmt.Printf("Enter the menu you want to order : ")
	utils.ReceiveUserInput(HandleFoodOrderInput)
}

func getFoodCategory() []string{
  var order []string
  foodMap := *food.GetFood()
  
  for key,_ := range foodMap{
    order = append(order, key)
  }
  sort.Strings(order)
  return order
}

func ShowFoodTable(){
  foodMap := *food.GetFood()
	number := 1
	result := ""

  foodMenuOrder := getFoodCategory()

	for _, category := range foodMenuOrder {
		var foodNameColumn []string
		var foodPriceColumn []string
		var foodStockColumn []string

		for _, food := range foodMap[category] {
            strNum := strconv.Itoa(number)
            if len(strNum) == 1 {
                strNum = " "+strNum
            }
            name := strNum + " " + strings.Title(food.Name)
            
			price := strconv.Itoa(food.Price) + "K"
			if len(price) == 1 {
				price = " " + price
			}
            
			stock := "stock : " + strconv.Itoa(food.Stock)

			foodNameColumn = append(foodNameColumn, name)
			foodPriceColumn = append(foodPriceColumn, price)
			foodStockColumn = append(foodStockColumn, stock)
			number++
		}
		result += utils.CreateTable(category, 50, 5, foodNameColumn, foodPriceColumn, foodStockColumn) + "\n"
	}
	fmt.Println(result)
}

func HandleFoodOrderInput(input string) {
  intInput, err := strconv.Atoi(input)
  foodMenuOrder := getFoodCategory()
  
  if err != nil {
    utils.SendMessage(nil, ShowFoodMenu)
    return
  }
  
  if intInput == 101 {
    ShowOrderMenu()
    return
  } else if intInput == 109 {
    ShowMainmenu()
  }
  
	var foodCount int
	foodMap := *food.GetFood()
	index := 0
	key := ""

	for _, k := range foodMenuOrder {
		index = intInput-1 - foodCount
		foodCount += len(foodMap[k])

		if foodCount >= intInput {
			key = k
			break
		}
	}
  
  if index > len(foodMap[key]) {
    utils.ClearScreen()
    utils.SendMessage("Food not found in menu !", ShowFoodMenu)
    return
  }
  
	SetupOrder(&foodMap[key][index])
}
