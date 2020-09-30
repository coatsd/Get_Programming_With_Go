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

type location struct {
	rOrL, site string
	lat, long coordinate
}

func newLocation(rOrL, site string, lat, long coordinate) location {
	return location{rOrL, site, lat, long}
}

func (l location) toDecimal() (string, string, float64, float64) {
	return l.rOrL, l.site, l.lat.toDecimal(), l.long.toDecimal()
}

func main() {
	locations := []location {
		{rOrL: "Spirit", site: "Columbia Memorial Station", 
			lat: coordinate{d: 14.0, m: 34.0, s: 6.2, h: 'S'}, 
			long: coordinate{d: 175.0, m: 28.0, s: 21.5, h: 'E'},
		},
	}

	for _, l := range locations {
		n, s, la, lo := l.toDecimal()
		fmt.Printf("%v - %v %v° %v°", n, s, la, lo)
	}
}