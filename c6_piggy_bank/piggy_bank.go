package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var piggy float32 = 0.0

	change := [3]float32 {
		0.05,
		0.10,
		0.25,
	}
	change_count := [3]int16 {
		0,
		0,
		0,
	}

	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	for piggy < 20.0 {
		change_index := rng.Intn(3)
		piggy += change[change_index]
		switch change_index {
		case 0:
			change_count[0] += 1
		case 1:
			change_count[1] += 1
		case 2:
			change_count[2] += 1
		}
	}

	fmt.Printf("Nickles: %v, Dimes: %v, Quarters: %v, Balance: $%v", 
		change_count[0], change_count[1], change_count[2], piggy)
}