package main

import "fmt"

// TODO: declare a type for Student (with first and last name)
type Student struct {
	FirstName string
	LastName  string
}

// TODO: declare a type for Class (consisting of multiple students)
type Class struct {
	Name     string
	Students []Student
}

// TODO: declare a map of modules being attended by multiple classes
var modules = map[string][]Class{
	"Math": {
		{Name: "Class A", Students: []Student{{"Kobert", "Runz"}, {"Nelio", "Gautschi"}}},
		{Name: "Class B", Students: []Student{{"Eric", "Hose"}, {"Travis", "Scott"}}},
	},
	"Science": {
		{Name: "Class C", Students: []Student{{"Aurel", "Schwitzer"}, {"Metro", "Boomin"}}},
		{Name: "Class D", Students: []Student{{"Stricher", "Bahnhofstricher"}, {"Kante", "B"}}},
	},
}

func main() {
	// Output everything using fmt.Println()
	fmt.Println("Modules and Classes:")
	for module, classes := range modules {
		fmt.Println(module)
		for _, class := range classes {
			fmt.Printf("- %s: Students - %v\n", class.Name, class.Students)
		}
		fmt.Println()
	}
}
