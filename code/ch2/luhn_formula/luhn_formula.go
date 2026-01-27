package main

import "fmt"

func doubleDigitValue(digit int) int {
	doubledDigit := digit * 2
	sum := 0
	if doubledDigit > 10 {
		sum = 1 + doubledDigit % 10
	} else {
		sum = doubledDigit
	}
	return sum
}

func main() {
	var b int = 0
	fmt.Scan(&b)
}