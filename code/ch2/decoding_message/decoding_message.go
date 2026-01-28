package main

import (
	"fmt"
	"strings"
)

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

func symbolTable(charArg int) string {
	var outputChar rune
	switch charArg {
		case 1:
			outputChar = '!'
		case 2:
			outputChar = '?'
		case 3:
			outputChar = ','
		case 4:
			outputChar = '.'
		case 5:
			outputChar = ' '
		case 6:
			outputChar = ';'
		case 7:
			outputChar = '"'
		case 8:
			outputChar = ','
	}
	return string(outputChar)
}

func main() {
	var getInput string
	var digitChar string
	var decodedMessage strings.Builder
	graph := map[string]string{
        "U": "L",        // A -> B
        "L": "P",        // B -> C  
        "P": "U",        // C -> A (creates the cycle)
    }
	var currentMode string = "U"
	fmt.Print("Enter your input : ")
	fmt.Scanln(&getInput)
	for i := 0; i < len(getInput); i++ {
		isLast := false
		if i == len(getInput) - 1 {
			digitChar = digitChar + string(getInput[i])
			isLast = true
		}
		if getInput[i] == ',' || isLast {
			getDigit := charToInt(digitChar)
			digitChar = ""
			if (currentMode == "U" || currentMode == "L") && getDigit % 27 == 0 {
				currentMode = graph[currentMode]
				continue
			}
			if currentMode == "P" && getDigit % 9 == 0 {
				currentMode = graph[currentMode]
				continue
			}
			switch currentMode {
				case "U":
					decodedMessage.WriteString(string((getDigit % 27) + 'A' - 1))
					continue
				case "L":
					decodedMessage.WriteString(string((getDigit % 27) + 'a' - 1))
					continue
				case "P":
					decodedMessage.WriteString(symbolTable(getDigit % 9))
					continue
			}
		} else {
			digitChar = digitChar + string(getInput[i])
		}
	}
	fmt.Println(decodedMessage.String())
}