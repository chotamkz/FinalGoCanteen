package view

import (
	"fmt"
	"main/model/finance"
	"main/utils"
	"strconv"
)

func ShowFinanceMenu(){
  financeDB := finance.GetFinanceDB()
  utils.ClearScreen()
  words := []string{"Total Income", "Total Expenditure"}
  finance := []string{utils.ToCurrency(financeDB.TotalIncome),
                     utils.ToCurrency(financeDB.TotalExpenditure)}

  menu := "\n\n[1] Show cash history"
  menu += "\n[2] Exit to main menu"
  menu += "\n\nEnter the menu you choose : "
  
  fmt.Println(utils.CreateTable("Finance Menu", 50, 4, words, finance))
  fmt.Println(utils.CreateLine(50,"="))
  fmt.Printf(menu)

  utils.ReceiveUserInput(handleFinanceMenu)
}

func handleFinanceMenu(input string){
  intInput, err := strconv.Atoi(input)
  if err != nil {
    utils.SendMessage(nil, ShowFinanceMenu)
  }

  switch intInput {
    case 1:
      utils.SendMessage("Feature not ready yet", ShowMainmenu)
    case 2:
      ShowMainmenu()
  }
}