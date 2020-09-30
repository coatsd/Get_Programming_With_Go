package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h rune
}

func (c coordinate) toDecimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S','W','s','w':
		sign = -1
	}
	return sign * (c.d + c.m / 60 + c.s / 3600)
}

func (c coordinate) toRadius() float64 {
	return c.toDecimal() * math.Pi / 180
}

type location struct {
	name string
	lat, long coordinate
}

type world struct {
	radius float64
}

func (w world) distance(l1, l2 location) float64 {
	s1, c1 := math.Sincos(l1.lat.toRadius())
	s2, c2 := math.Sincos(l2.lat.toRadius())
	clong := math.Cos(l1.long.toRadius() - l2.long.toRadius())
	return w.radius * math.Acos(s1 * s2 + c1 * c2 * clong)
}

func (w world) printDis(l1, l2 location) {
	fmt.Printf("the distance between %v and %v is %v\n", l1.name, l2.name, w.distance(l1, l2))
}

func runMarsTests() {
	mars := world{3389.5}
	marsSites := []location{
		location{
			name: "spirit",
			lat: coordinate{d: 14.0, m: 34.0, s: 6.2, h: 'S',}, 
			long: coordinate{d: 175.0, m: 28.0, s: 21.5, h: 'E',},
		},
		location{
			name: "opportunity",
			lat: coordinate{d: 1.0, m: 56.0, s: 46.3, h: 'S',},
			long: coordinate{d: 354.0, m: 28.0, s: 24.2, h: 'E',},
		},
		location{
			name: "curiosity",
			lat: coordinate{d: 4.0, m: 35.0, s: 22.2, h: 'S',},
			long: coordinate{d: 137.0, m: 26.0, s: 30.1, h: 'E',},
		},
		location{
			name: "insight",
			lat: coordinate{d: 4.0, m: 35.0, s: 0.0, h: 'N',},
			long: coordinate{d: 135.0, m: 54.0, s: 0.0, h: 'E',},
		},
	}

	fmt.Println("Mars Tests:")
	for i, l1 := range marsSites {
		for j, l2 := range marsSites {
			if i < j {
				mars.printDis(l1, l2)
			}
		}
	}
}

func runEarthTests() {
	earth := world{6371.0}
	london := location{
		name: "london",
		lat: coordinate{d: 51.0, m: 30.0, h: 'N',},
		long: coordinate{d: 0.0, m: 8.0, h: 'W',},
	}
	paris := location{
		name: "paris",
		lat: coordinate{d: 48.0, m: 51.0, h: 'N',},
		long: coordinate{d: 2.0, m: 21.0, h: 'E',},
	}

	mtSharp := location{
		name: "mount sharp",
		lat: coordinate{d: 5.0, m: 4.0, s: 48.0, h: 'S',},
		long: coordinate{d: 137.0, m: 51.0, h: 'E',},
	}
	Olymp := location{
		name: "Olympus Mons",
		lat: coordinate{d: 18.0, m: 39.0, h: 'N'},
		long: coordinate{d: 226.0, m: 12.0, h: 'E',},
	}

	fmt.Println("Earth Tests:")
	earth.printDis(london, paris)
	earth.printDis(mtSharp, Olymp)
}

func main() {
	fmt.Printf("\n")
	runMarsTests()
	fmt.Printf("\n")
	runEarthTests()
	fmt.Printf("\n")
}