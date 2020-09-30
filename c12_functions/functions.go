package main

import (
	"fmt"
)

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToKelvin(c float64) float64 {
	c += 273.15
	return c
}

func celsiusToFahren(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func fahrenToCelsius(f float64) float64 {
	return (f - 32.0) * 5.0 / 9.0
}

func fahrenToKelvin(f float64) float64 {
	return celsiusToKelvin(fahrenToCelsius(f))
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	fahren := celsiusToFahren(celsius)
	fmt.Printf("%v kelvin is %v celsius\n", kelvin, celsius)
	fmt.Printf("%v celsius is %v Fahrenheit\n", celsius, fahren)
	kelvin = fahrenToKelvin(fahren)
	fmt.Printf("%v fahrenheit is %v kelvin\n", fahren, kelvin)
}