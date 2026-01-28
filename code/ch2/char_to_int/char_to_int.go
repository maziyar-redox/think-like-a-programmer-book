package main

import "fmt"

func main() {
	var charArg string = "48"
	var sum int
	for i := 0; i < len(charArg); i++ {
		sum += int(charArg[i] - 48)
		if i <= len(charArg) {
			sum *= 10
		}
	}
	sum /= 10
	fmt.Println(sum)
}