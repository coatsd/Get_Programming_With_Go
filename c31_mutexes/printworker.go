package main

import (
	"fmt"
	"sync"
	"time"
)

func printWorker() {
	n := 0
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			n++
			fmt.Println(n)
			next = time.After(time.Second)
		}
	}
}

func main() {
	
}