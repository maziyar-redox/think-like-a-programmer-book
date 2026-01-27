package main

import "fmt"


func main() {
	var nums int
	var positiveNums int
	var negativeNums int
	for i := 0; i < 10; i++ {
		fmt.Scanln(&nums)
		if nums < 0 {
			negativeNums++
			continue
		} else {
			positiveNums++
			continue
		}
	}
	fmt.Println("Amount of negative numbers : ", negativeNums)
	fmt.Println("Amount of positive numbers : ", positiveNums)
}