package main

import (
	"fmt"
	"time"
	"flag"
	"math/rand"
)

func main() {
	numPt := flag.Int("v", 2, "Run a particular listing version")

	flag.Parse()

	switch *numPt {
	case 1:
		verOne()
	case 2:
		verTwo()
	case 3:
		verThree()
	default:
		verThree()
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
		// When "sleepyGopher()" hits 
		gopherId := <-c
		fmt.Printf("gopher %v has finished sleeping\n", gopherId)
	}
}

// code from listing 30.5 describes use of select statements, which run when a
// value is received in one of the channels in a select case. It also describes
// use of the time.After method to set a timeout.
func verThree() {
	sleepyGopher := func(id int, c chan int) {
		time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
		c <- id
	}

	c := make(chan int)
	timeout := time.After(3 * time.Second)

	fmt.Println("Running code from listing 30.5")

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c:
			fmt.Printf("gopher %v has finished sleeping\n", gopherId)
		case <-timeout:
			fmt.Println("my patience has ran out")
			return
		}
	}
}