package main

import (
	"fmt"
)

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(tempCelsius float64) float64 {
	tempFahrenheit := tempCelsius*(9.0/5.0) + 32.0
	return tempFahrenheit
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(tempFahrenheit float64) float64 {
	tempCelsius := (tempFahrenheit - 32.0) * (5.0 / 9.0)
	return tempCelsius
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	temp1 := convertCelsiusToFahrenheit(20.0)
	fmt.Println(temp1) // 68

	temp2 := convertCelsiusToFahrenheit(15.0)
	fmt.Println(temp2) // 59

	temp3 := convertCelsiusToFahrenheit(30.0)
	fmt.Println(temp3) // 86

	fmt.Println("______________________________________")

	// TODO: call the function convertFahrenheitToCelsius
	temp4 := convertFahrenheitToCelsius(68.0)
	fmt.Println(temp4) // 20

	temp5 := convertFahrenheitToCelsius(59.0)
	fmt.Println(temp5) // 15

	temp6 := convertFahrenheitToCelsius(86.0)
	fmt.Println(temp6) // 30
}
