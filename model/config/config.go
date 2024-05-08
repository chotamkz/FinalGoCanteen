package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
  Password string
}

var PASSWORD string
var m_user User

func Init(){
  data, err := ioutil.ReadFile("DB/user.json")
  if err != nil {
    log.Fatal(err)
  }
  json.Unmarshal(data, &m_user)
  PASSWORD = m_user.Password
}

func SetNewPassword(password string){
  m_user.Password = password
  data, err := json.Marshal(m_user)
  
  if err != nil {
    log.Fatal(err)
  }

  os.WriteFile("DB/user.json", data, 0666)
}