package view

import (
	"fmt"
	"main/model/finance"
	"main/model/order"
	"main/utils"
	"strconv"
)

func ShowPaymentMenu(bill int){
  moneyPaid := 0
  fmt.Printf("\nEnter amount of money : ")
  fmt.Scanln(&moneyPaid)

  showInvoice(bill, moneyPaid)
}

func showInvoice(bill int, moneyPaid int) {
  utils.ClearScreen()
  finance.AddIncome(bill)
  
  orderTable := ShowOrderTable("Invoice")
  fmt.Println(orderTable)
  fmt.Println(utils.CreateLine(50,"-")+"\n")

  createLeftAlignText("Bill " + strconv.Itoa(bill))
  createLeftAlignText("Paid " + strconv.Itoa(moneyPaid))
  fmt.Println("")
  
  change := moneyPaid - bill
  if change != 0 {
    createLeftAlignText("Remain: " + strconv.Itoa(change))
  }
  
  fmt.Println("\n"+utils.CreateLine(50, "="))

  fmt.Println("\nThanks Thanks, come back in SDU Canteen Menu \n\nSDU")
  fmt.Printf("\nEnter to return to the main menu")
  fmt.Scanln()
  order.ClearOrder()
  ShowMainmenu()
}

func createLeftAlignText(text string){
  fmt.Printf(utils.CreateLine(50 - len(text), " "))
  fmt.Println(text)
}