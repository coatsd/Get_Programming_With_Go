package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var piggy int16 = 0

	change := [3]int16 {
		5,
		10,
		25,
	}

	change_count := [3]int16 {
		0,
		0,
		0,
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for piggy < 2000 {
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
		change_count[0], change_count[1], change_count[2], float32(piggy) / 100)
}