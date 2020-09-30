package main

import (
	"fmt"
)

func main() {
	message := "Hola Estación Espacial Internacional"
	for _, c := range message {
		if c < 'A' || c > 'z'  {
			fmt.Printf("%c", c)
		} else if c > 'Z' && c < 'n' || c < 'Z' && c < 'N' {
			fmt.Printf("%c", c + 13)
		} else {
			fmt.Printf("%c", c - 13)
		}
	}
}