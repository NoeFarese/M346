package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = fibs[1] + fibs[2]
	fibs[4] = fibs[2] + fibs[3]

	for i := 5; i < 9; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])
	}

	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
