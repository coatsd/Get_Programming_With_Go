package main

import (
	"fmt"
	"time"
	"flag"
)

func main() {
	numPt := flag.Int("v", 2, "Run a particular listing")

	flag.Parse()

	switch *numPt {
	case 1:
		verOne()
	case 2:
		verTwo()
	default:
		verTwo()
	}
}

// code from listing 30.1-30.3 describes use of goroutines.
func verOne() {
	sleepyGopher := func(id int) {
		// sleep, to demonstrate the overhead of a routine that takes time. 
		time.Sleep(3 * time.Second)
		fmt.Printf("...%v snore...\n", id)
	}

	fmt.Println("Running code from listing 30.1-30.3")

	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	// Waiting 4 seconds for the goroutines to finish - obviously not the best.
	time.Sleep(4 * time.Second)
}

// code from listing 30.4 describes implementing channels to give the code a
// way to keep track of when a goroutine finishes, so that the programmer
// doesn't have to guess or use non-idiomatic methods of checking.
func verTwo() {
	sleepyGopher := func(id int, c chan int) {
		time.Sleep(3 * time.Second)
		fmt.Printf("...%v snore...\n", id)
		c <- id
	}

	c := make(chan int)

	fmt.Println("Running code from listing 30.4")

	// This loop spins up the goroutines.
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	// This loop waits for the channel to send an integer from each goroutine call
	for i := 0; i < 5; i++ {
		gopherId := <-c
		fmt.Printf("gopher %v has finished sleeping\n", gopherId)
	}
}