package main

import "fmt"

func charToInt(charArg string) int {
	var sum int
	for i := 0; i < len(charArg); i++ {
		sum += int(charArg[i] - 48)
		if i <= len(charArg) {
			sum *= 10
		}
	}
	sum /= 10
	return sum
}

func main() {
	var getInput string
	var digitChar string
	//var decodedMessage string
	fmt.Print("Enter your input : ")
	fmt.Scanln(&getInput)
	for i := 0; i < len(getInput); i++ {
		if getInput[i] == ',' || getInput[i] == 0 {
			fmt.Println("First int is : ", digitChar, "Which is : ", charToInt(digitChar))
		} else {
			digitChar = digitChar + string(getInput[i])
		}
	}
}