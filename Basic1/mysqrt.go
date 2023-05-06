package main


import (
	"fmt"
	"math"
)

// http://composingprograms.com/pages/16-higher-order-functions.html
func Sqrt(x float64) float64 {
	guess := 1.0
	for math.Abs(x - guess * guess) > 1E-15 { // 53log10(2) = 15.954589770191003
		fmt.Println(guess)
		guess -= (guess*guess - x) / (2*guess)
	}
	return guess
}



func approx_eq(x, y float64) bool {
	return math.Abs(x - y) < 1E-15
}

func near_zero(x float64) bool {
	return approx_eq(x, 0)
}

func main() {
	fmt.Println(Sqrt(4))
}
