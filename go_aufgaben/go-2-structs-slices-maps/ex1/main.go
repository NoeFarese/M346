package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   byte
	Month byte
	Year  int16
}

type Profile struct {
	Name             FullName
	Born             BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	myName := FullName{
		FirstName: "Noé",
		LastName:  "Farese",
	}

	myBirthDate := BirthDate{
		Day:   24,
		Month: 9,
		Year:  2006,
	}

	me := Profile{
		Name:             myName,
		Born:             myBirthDate,
		NumberOfSiblings: 1,
		ZodiacSign:       'L',
	}

	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)

	me.NumberOfSiblings++

	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
