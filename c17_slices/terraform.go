package main

import (
	"fmt"
)

type planets []string

func terraform(ps planets) []string {
	for i := range ps {
		ps[i] = "New " + ps[i]
	}
	return ps
}

func main() {
	planets := [8]string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	pSlice := terraform(planets[:])

	fmt.Printf("Planets:\n")
	for i := range pSlice {
		fmt.Printf(pSlice[i] + "\n")
	}
}