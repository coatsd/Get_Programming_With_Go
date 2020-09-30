package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const airline_options = 3
	const trip_type_options = 2
	const min_days = 14
	const day_var = 15
	const min_price = 50
	const price_var = 51

	airlines := [airline_options]string {
		"Virgin Galactic",
		"SpaceX",
		"Space Adventures",
	}
	trip_types := [trip_type_options]string {
		"One-Way",
		"Round-Trip",
	}

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	fmt.Printf("%-20v %-5v %-15v %5v\n", "Airline", "Days", "Trip Type", "Cost")
	fmt.Printf("================================================\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("%-20v %-5v %-15v $%4v\n",
			airlines[rng.Intn(airline_options)],
			rng.Intn(day_var) + min_days,
			trip_types[rng.Intn(trip_type_options)],
			rng.Intn(price_var) + min_price)
	}
}