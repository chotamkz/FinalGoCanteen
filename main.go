package main

import (
	"main/model/finance"
	"main/model/food"
	"main/view"
)

func main() {
	food.Init()
	finance.Init()
	view.ShowMainmenu()
}
