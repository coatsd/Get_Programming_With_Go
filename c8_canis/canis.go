package main

import (
	"fmt"
)

func main() {
	const (
		lightSpeed = 299792
		secondsPerDay = 86400
		daysPerYear = 365
		distance = 236000000000000000
		days = distance / lightSpeed / secondsPerDay / daysPerYear
	)
	fmt.Println("Canis Major is ", days, " light days away")
}