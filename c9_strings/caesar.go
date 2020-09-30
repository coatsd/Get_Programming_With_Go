package main

import (
	"fmt"
)

func main() {
	var message string = "L fdph, L vdz, L frqtxhuhg"
	for _, c := range message {
		if c < 'A' || c > 'z'  {
			fmt.Printf("%c", c)
		} else if c < 'd' && c > 'Z' {
			fmt.Printf("%c", c + 23)
		} else {
			fmt.Printf("%c", c - 3)
		}
	}
}