package utils

import (
	"fmt"
	"strconv"
)

func CreateLine(count int, char string) string {
	result := ""
	for i := 0; i < count; i++ {
		result += char
	}
	return result
}

func GetLongestWord(words []string) int {
	max := 0
	for _, word := range words {
		if len(word) > max {
			max = len(word)
		}
	}
	return max
}

func ClearScreen() {
	fmt.Print("\x1bc")
}

func ReceiveUserInput(inputHandle func(input string)) {
	userInput := ""
	fmt.Scanln(&userInput)

	inputHandle(userInput)
}

func SendMessage(msg interface{}, showFunction func()) {
	ClearScreen()
	var ans string
	var invalidMessage string

	if msg == nil {
		invalidMessage = "Invalid Input"
	} else {
		invalidMessage = msg.(string)
	}

	invalidMessage += "\n\nPress Enter to Continue..."

	fmt.Printf(invalidMessage)
	fmt.Scanln(&ans)
	showFunction()
}

func ToCurrency(value int) string {
	strValue := strconv.Itoa(value)
	for i := len(strValue) - 3; i > 0; i -= 3 {
		strValue = strValue[:i] + "." + strValue[i:]
	}

	return strValue
}
