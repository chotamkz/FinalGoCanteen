package view

import (
	"fmt"
	"main/model/config"
	"main/model/food"
	"main/utils"
	"strconv"
	"strings"
)

func ShowInputPasswordForm() {
  config.Init()
	var password string

	fmt.Printf("Input a password : ")
	fmt.Scanln(&password)

	passwordValidation(password)
}

func passwordValidation(password string) {
	if password != config.PASSWORD {
		utils.SendMessage("Wrong credentials", ShowMainmenu)
		return
	}

	showConfigMenu()
}

func showConfigMenu() {
	utils.ClearScreen()
	menu := utils.CreateTitle(50, "Configuration")
	menu += "\n\n[1] Add new menu"
	menu += "\n[2] Remove menu"
	menu += "\n[3] Change password"
  menu += "\n[4] Back to main menu\n"

	fmt.Println(menu)
	fmt.Printf("Enter the menu you choose : ")
	utils.ReceiveUserInput(handleConfigMenu)
}

func handleConfigMenu(input string) {
	intInput, err := strconv.Atoi(input)
	if err != nil {
		utils.SendMessage(nil, showConfigMenu)
	}

	switch intInput {
	case 1:
		addNewFood()
  case 2:
    removeFoodMenu()
	case 3:
		changePassword()
  case 4:
    ShowMainmenu()
	default:
		utils.SendMessage(nil, showConfigMenu)
	}
}

func addNewFood() {
	utils.ClearScreen()
	var category, foodName string
	var foodPrice, foodStock int

	fmt.Printf("Input food category : ")
	fmt.Scanln(&category)
  
	fmt.Printf("Input food name : ")
	fmt.Scanln(&foodName)
  
	fmt.Printf("Input food price (35.000 -> 35) : ")
	fmt.Scanln(&foodPrice)
  
	fmt.Printf("Input food stock : ")
	fmt.Scanln(&foodStock)
  
	newFood := food.Food{
		Name:  strings.ToLower(foodName),
		Price: foodPrice,
		Stock: foodStock,
	}

	food.Add(strings.ToLower(category), newFood)
  utils.SendMessage("New food success added !", ShowMainmenu)
}

func removeFoodMenu(){
  utils.ClearScreen()
  var category, foodName string
  foodMap := *food.GetFood()

  fmt.Printf("Input food category : ")
  fmt.Scanln(&category)

  if _,ok := foodMap[strings.ToLower(category)]; ok == false {
    utils.SendMessage("Food category not found !", showConfigMenu)
    return
  }

  fmt.Printf("Input food name : ")
  fmt.Scanln(&foodName)

  foodIndex := foodMap[category].FindIndex(strings.ToLower(foodName))
  if foodIndex == -1 {
    utils.SendMessage("Food not found !", showConfigMenu)
    return
  }

  food.Remove(category, foodIndex)
  utils.SendMessage("Food has been removed", ShowMainmenu)
}

func changePassword() {
	utils.ClearScreen()
	var oldPassword, newPassword string

	fmt.Printf("Input the old password : ")
	fmt.Scanln(&oldPassword)

	if oldPassword != config.PASSWORD {
		utils.SendMessage("Invalid Password", showConfigMenu)
		return
	}

	fmt.Printf("Input a new password : ")
	fmt.Scanln(&newPassword)

	config.SetNewPassword(newPassword)
	utils.SendMessage("Password changed !", ShowMainmenu)
}
