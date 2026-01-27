package main

import "fmt"

func doublingDigitValue(digit int) int {
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
	var userInput int
	var oddLengthChecksum int
	var evenLengthChecksum int
	fmt.Scanln(&userInput)
	position := 1
	for userInput % 10 != 0 {
		currentNumber := userInput % 10
		if position % 2 == 0 {
			oddLengthChecksum += doublingDigitValue(currentNumber)
			evenLengthChecksum += currentNumber
		} else {
			oddLengthChecksum += currentNumber
			evenLengthChecksum += doublingDigitValue(currentNumber)
		}
		userInput /= 10
		position++
	}
	var checksum int
	if (position - 1) % 2 == 0 {
		checksum = evenLengthChecksum
	} else {
		checksum = oddLengthChecksum
	}
	fmt.Println("Checksum is :", checksum)
	if checksum % 10 == 0 {
		fmt.Println("Checksum is valid")
	} else {
		fmt.Println("Checksum is not valid")
	}
}