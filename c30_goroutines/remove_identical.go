package main

import (
	"fmt"
)

func source(ds chan string) {
	for _, v := range []string{"value 1", "value 1", "value 2", "value 3"} {
		ds <- v
	}
	close(ds)
}

func filter(us, ds chan string) {
	var prev string
	for {
		v, ok := <-us
		if !ok {
			break
		}
		if v != prev {
			ds <- v
		}
		prev = v
	}
	close(ds)
}

func printString(us chan string) {
	for {
		v, ok := <-us
		if !ok {
			return
		}
		fmt.Println(v)
	}
}

func runCompose() {
	c1 := make(chan string)
	c2 := make(chan string)
	go source(c1)
	go filter(c1,c2)
	printString(c2)
}

func main() {
	runCompose()
}