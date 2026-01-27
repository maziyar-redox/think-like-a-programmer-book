package main

import (
	"fmt"
	"math"
)

func main() {
	width := 7
	height := math.Abs(float64(width/2)) + 1
	for i := 1; i <= width; i++ {
		for j := 1; j <= int(height - math.Abs(height - float64(i))); j++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}