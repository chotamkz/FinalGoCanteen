package food

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Food struct {
	Name         string
	Price, Stock int
}

type foodList []Food
type foodMap map[string] foodList

var foods foodMap
var FoodMenuOrder []string

func Init() {
	data, err := ioutil.ReadFile("DB/foodData.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(data, &foods)
}

func GetFood() *foodMap {
	return &foods
}

func Add(category string, data Food) {
	foods[category] = append(foods[category], data)
  save()
}

func Remove(category string, foodIndex int){
  foods[category] = append(
    foods[category][:foodIndex], 
    foods[category][foodIndex+1:]...
  )

  if len(foods[category]) == 0 {
    delete(foods, category)
  }

  save()
}

func UpdateStock(food *Food, value int){
  food.Stock -= value
  save()
}

func (foods foodList) FindIndex(name string) int {
  for index, value := range foods{
    if value.Name == name {
      return index
    }
  }
  
  return -1
}

func save(){
  data, err := json.Marshal(foods)
  if err != nil {
    log.Fatal(err)
  }

  os.WriteFile("DB/foodData.json", data, 0666)
}