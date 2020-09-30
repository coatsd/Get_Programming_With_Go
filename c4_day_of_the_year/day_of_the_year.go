package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	src := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(src)
	year := gen.Intn(time.Now().Year())
	month := gen.Intn(12) + 1
	days_in_month := 31

	switch month {
	case 2:
		if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
			days_in_month = 29
		} else {
			days_in_month = 28
		}
	case 4, 6, 9, 11:
		days_in_month = 30
	}

	day := gen.Intn(days_in_month) + 1
	fmt.Printf("Month: %v, Day: %v, Year: %v", month, day, year)
}