package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(100) + 1
	var guess = 0

	for num != guess {
		guess = rand.Intn(100) + 1
		if guess < num {
			fmt.Printf("%v is too low!\n", guess)
		} else if guess > num {
			fmt.Printf("%v is too high!\n", guess)
		} else {
			fmt.Printf("%v is correct!\n", guess)
		}
	}
}