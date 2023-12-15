package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a, b, c float64) []float64 {
	d := math.Pow(b, 2) - 4*a*c

	var solutions []float64

	if d > 0 {
		solution1 := (-b + math.Sqrt(d)) / (2 * a)
		solution2 := (-b - math.Sqrt(d)) / (2 * a)
		solutions = append(solutions, solution1, solution2)
	} else if d == 0 {
		solution := -b / (2 * a)
		solutions = append(solutions, solution)
	}

	return solutions
}

func main() {
	// TODO: call the function computeQuadraticFormula
	result1 := computeQuadraticFormula(3, 4, 1)
	fmt.Println(result1) // [-0.3333333333333333 -1]

	result2 := computeQuadraticFormula(2, 4, 2)
	fmt.Println(result2) // [-1]

	result3 := computeQuadraticFormula(3, 4, 2)
	fmt.Println(result3) // []
}
