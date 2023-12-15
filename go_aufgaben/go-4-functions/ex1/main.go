package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints float64, maxPoints float64) float64 {
	grade := gotPoints/maxPoints*5 + 1

	if grade < 1.0 {
		return 1.0
	} else if grade > 6.0 {
		return 6.0
	}

	return grade
}

func main() {
	// TODO: call the function computeGrade
	grade1 := computeGrade(17.5, 28.0)
	fmt.Println("Grade 1:", grade1) // 4.125

	grade2 := computeGrade(20.0, 30.0)
	fmt.Println("Grade 2:", grade2) // 4.333333333333333

	grade3 := computeGrade(15.0, 25.0)
	fmt.Println("Grade 3:", grade3) // 4
}
