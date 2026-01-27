package main

import "fmt"

func main() {
	dim := 4
	for i := 1; i <= dim; i++ {
		for j := 1; j <= dim + 1 - i; j++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}