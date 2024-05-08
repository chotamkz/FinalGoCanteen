package utils

import (
  "reflect"
  "fmt"
)

func CreateTable(title interface{}, titleLength int, columnSpace int, columnContents ...interface{}) string {
  var content []string
  response := ""
  
  if title != nil {
    strTitle := fmt.Sprint(title)
    response += CreateTitle(titleLength, strTitle)
    response += "\n\n"
  }

  for colIndex, columnContent := range columnContents{
    sliceContent := columnContentConvert(columnContent)
    longestWord := GetLongestWord(sliceContent)
    
    if content == nil {
      content = make([]string, len(sliceContent))
    }
    
    for i, c := range sliceContent{
      space := longestWord + columnSpace - len(c)
      if colIndex == len(columnContents)-1 {
        space = 0
      }
      content[i] += c + CreateLine(space , " ") 
    }
  }

  for _, c := range content {
    response += c + "\n"
  }
  return response
}

func columnContentConvert(a interface{}) []string{
  refContent := reflect.ValueOf(a)
  result := []string{}
  for i:=0; i<refContent.Len(); i++{
    strContent := fmt.Sprint(refContent.Index(i).Interface())
    result = append(result, strContent)
  }
  return result
}