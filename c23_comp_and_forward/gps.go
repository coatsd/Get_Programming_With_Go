package main

import (
	"fmt"
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

type gps struct {
	curr, dest location
	w world
}

func main() {
	
}