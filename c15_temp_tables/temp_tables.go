package main

import (
	"fmt"
)

type fahrenheit float64
type celsius float64
const p = "|"
const b = "=================================\n"
const template = "%-1v %-6v %-1v %-20v %-1v\n"

func celsiusToFahren(c celsius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func fahrenToCelsius(f fahrenheit) celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func genCtoF(beginTemp celsius) {
	fmt.Printf(b)
	fmt.Printf(template, p, "°C", p, "°F", p)
	fmt.Printf(b)
	for t := beginTemp; t <= 100.0; t = t + 5 {
		fmt.Printf(template, p, t, p, celsiusToFahren(t), p)
	}
	fmt.Printf(b)
}

func genFtoC(beginTemp fahrenheit) {
	fmt.Printf(b)
	fmt.Printf(template, p, "°F", p, "°C", p)
	fmt.Printf(b)
	for t := beginTemp; t <= 100.0; t = t + 5 {
		fmt.Printf(template, p, t, p, fahrenToCelsius(t), p)	
	}
	fmt.Printf(b)
}

func main() {
	const beginTemp float64 = -40.0
	genCtoF(celsius(beginTemp))
	genFtoC(fahrenheit(beginTemp))
}