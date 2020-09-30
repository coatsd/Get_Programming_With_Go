package main

import (
	"fmt"
)

func main() {
	// constants were omitted, because they're unnecessary for this
	fmt.Printf("The MPH required to reach Mars in %v days is %v MPH", 
		28, 28.0 * 24.0 / 56000000.0)
}