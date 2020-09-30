package main

import (
	"fmt"
)

type kelvin float64
type fahrenheit float64
type celsius float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func celsiusToFahren(c celsius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func fahrenToCelsius(f fahrenheit) celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func fahrenToKelvin(f fahrenheit) kelvin {
	return celsiusToKelvin(fahrenToCelsius(f))
}

func kelvinToFahren(k kelvin) fahrenheit {
	return celsiusToFahren(kelvinToCelsius(k))
}

func main() {
	var k kelvin = 233.0
	c := kelvinToCelsius(k)
	f := celsiusToFahren(c)
	fmt.Printf("%v kelvin is %v celsius\n", k, c)
	fmt.Printf("%v celsius is %v Fahrenheit\n", c, f)
	k = fahrenToKelvin(f)
	fmt.Printf("%v fahrenheit is %v kelvin\n", f, k)
}