package view

import (
	"fmt"
	"main/utils"
	"os"
	"strconv"
)

func ShowMainmenu() {
	output := ""
	output += utils.CreateTitle(30, "SDU Canteen Menu")
	output += "\n\n"
	output += "[1] Order\n"
	output += "[2] Finance\n"
	output += "[3] Managment order\n"
	output += "[4] Exit\n\n"
	output += "Enter the menu you choose : "

	utils.ClearScreen()
	fmt.Printf(output)
	utils.ReceiveUserInput(parsingInputToInt)
}

func parsingInputToInt(input string) {
	if intInput, err := strconv.Atoi(input); err == nil {
		handleMainmenu(intInput)
		return
	}

	utils.SendMessage(nil, ShowMainmenu)
}

func handleMainmenu(input int) {
	switch input {
	case 1:
		ShowFoodMenu()
	case 2:
		ShowFinanceMenu()
	case 3:
    ShowInputPasswordForm()
	case 4:
		os.Exit(1)
	default:
		utils.SendMessage(nil, ShowMainmenu)
	}
}
