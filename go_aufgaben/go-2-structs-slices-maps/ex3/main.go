package main

import "fmt"

func main() {
	modules := make(map[int]string)

	modules[104] = "Modul 104"
	modules[117] = "Modul 117"
	modules[346] = "Modul 346"

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)
	modules[208] = "Modul 208"
	modules[346] = "M346"

	fmt.Println(modules)
}
