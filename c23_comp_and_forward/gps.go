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

type gps struct {
	curr, dest location
	w world
}

func (x gps) distance() float64 {
	s1, c1 := math.Sincos(x.curr.lat.toRadius())
	s2, c2 := math.Sincos(x.dest.lat.toRadius())
	clong := math.Cos(x.curr.long.toRadius() - x.dest.long.toRadius())
	return x.w.radius * math.Acos(s1 * s2 + c1 * c2 * clong)
}

func (x gps) printDis() {
	fmt.Printf("the distance between %v and %v is %v kilometers\n", x.curr.name, x.dest.name, x.distance())
}

type rover struct {
	name string
	gps
}

func (r rover) message() {
	r.printDis()
}

func gpsTests() {
	gpsTest := gps{
		curr: location{
			name: "you",
			lat: coordinate{d: 51.0, m: 30.0, h: 'N',},
			long: coordinate{d: 0.0, m: 8.0, h: 'W',},
		},
		dest: location{
			name: "paris",
			lat: coordinate{d: 48.0, m: 51.0, h: 'N',},
			long: coordinate{d: 2.0, m: 21.0, h: 'E',},
		},
		w: world{6371.0},
	}

	fmt.Println("GPS Tests:")
	// because gps is embedded, it forwards the printDis call to the embedded
	// gps object.
	gpsTest.printDis()
}

func roverTests() {
	r := rover{
		"curiosity",
		gps{
			curr: location{
				name: "bradbury landing",
				lat: coordinate{d: 4.0, m: 35.0, s: 22.2, h: 'S',},
				long: coordinate{d: 137.0, m: 26.0, s: 30.1, h: 'E',},
			},
			dest: location{
				name: "elysium planitia",
				lat: coordinate{d: 4.0, m: 35.0, s: 0.0, h: 'N',},
				long: coordinate{d: 135.0, m: 54.0, s: 0.0, h: 'E',},
			},
			w: world{3389.5},
		},
	}

	fmt.Println("Rover Tests:")
	r.message()
}

func main() {
	gpsTests()
	fmt.Printf("\n")
	roverTests()
}