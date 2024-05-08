package finance

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type financeModel struct {
  TotalIncome, TotalExpenditure int
}

var financeDB financeModel

func GetFinanceDB() financeModel{
  return financeDB
}

func AddIncome(value int){
  financeDB.TotalIncome += value
  save()
}

func Init(){
  data, err := ioutil.ReadFile("DB/financeData.json")
  if err != nil {
    log.Fatal(err)
  }

  json.Unmarshal(data, &financeDB)
}

func save(){
  json, err := json.Marshal(financeDB)
  if(err != nil){
    log.Fatal(err)
  }

  os.WriteFile("DB/financeData.json", json, 0666)
}