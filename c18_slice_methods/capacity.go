package main

import (
	"fmt"
)

func appendToSlice(intSlice *[]int, print bool) {
	if print {
		fmt.Printf("Adding ten items\n")
	}
	for i := 0; i < 10; i++ {
		*intSlice = append(*intSlice, i)
		if print {
			fmt.Printf("capacity: %v, array: %v\n", cap(*intSlice), *intSlice)
		}
	}
}

func main() {
	var sliceVar []int
	fmt.Printf("Test capacity of raw slice being appended to:\n")
	appendToSlice(&sliceVar, true)

	sliceVar = make([]int, 0, 5)
	fmt.Printf("Test capacity of slice with starting fixed capacity of 5:\n")
	appendToSlice(&sliceVar, true)

	sliceVar = make([]int, 0, 1)
	sliceVar2 := make([]int, 0, 1)
	appendToSlice(&sliceVar2, false)
	fmt.Printf("Testing capacity of slice using 1 call to append for 10 items:\n")
	sliceVar = append(sliceVar, sliceVar2...)
	fmt.Printf("capacity: %v, array: %v\n", cap(sliceVar), sliceVar)
}