package main

import (
	"fmt"
	"time"
	"flag"
	"math/rand"
	"strings"
)

func main() {
	// creates a flag for command line args, which determines what code to run.
	// The flag is -v, the default value is 0, and the last is a description of
	// the flag itself.
	numPt := flag.Int("v", 0, "Run a particular listing version")

	flag.Parse()

	switch *numPt {
	case 1:
		verOne()
	case 2:
		verTwo()
	case 3:
		verThree()
	case 4:
		verFour()
	default:
		verFive()
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

// code from listing 30.5-30.6 describes use of select statements, which run
// when a value is received in one of the channels in a select case. It also
// describes use of the time.After method to set a timeout.
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

// code from listing 30.7-30.10 describes use of channels in multiple functions.
// These functions are called "pipelines", which process data coming from other
// sources. verFour uses a "sentinel value", or zero value of the datatype to
// indicate the end of the stream.
func verFour() {
	// Pass a string "downstream" (ds)
	sourceGopher := func(ds chan string) {
		for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
			ds <- v
		}
		ds <- ""
	}
	// Receives a string from "upstream" (us), processes out "bad" strings, then
	// continues to pass other items "downstream"
	filterGopher := func(us, ds chan string) {
		for {
			item := <-us
			if item == "" {
				ds <- item
				return
			}
			if !strings.Contains(item, "bad") {
				ds <- item
			}
		}
	}
	// Receives an item from upstream and prints the item.
	printGopher := func(us chan string) {
		for {
			v := <-us
			if v == "" {
				return
			}
			fmt.Println(v)
		}
	}

	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c1)
	go filterGopher(c1, c2)
	printGopher(c2)
}

// code from listing 30.11-30.14 describes use of the close function to close a
// channel when the channel is no longer in use. If a channel is expected to be
// closed at any point during the program, you need to check if the channel has
// closed. This is done with the second value returned from a channel receiver.
// Example: value, isClosed := <-channel
func verFive() {
	// Pass a string "downstream" (ds)
	sourceGopher := func(ds chan string) {
		for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
			ds <- v
		}
		close(ds)
	}
	// Receives a string from "upstream" (us), processes out "bad" strings, then
	// continues to pass other items "downstream"
	filterGopher := func(us, ds chan string) {
		for {
			item, ok := <-us
			if !ok {
				close(ds)
				return
			}
			if !strings.Contains(item, "bad") {
				ds <- item
			}
		}
		close(ds)
	}
	// Receives an item from upstream and prints the item.
	printGopher := func(us chan string) {
		for v := range us {
			fmt.Println(v)
		}
	}

	c1 := make(chan string)
	c2 := make(chan string)
	go sourceGopher(c1)
	go filterGopher(c1, c2)
	printGopher(c2)
}