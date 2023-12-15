package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	// TODO: cann the function computeHypotenuse
	c1 := computeHypotenuse(15, 20)
	fmt.Println(c1) // 25

	c2 := computeHypotenuse(420, 36)
	fmt.Println(c2) // 421.5400336860071

	c3 := computeHypotenuse(12, 5)
	fmt.Println(c3) // 13
}
