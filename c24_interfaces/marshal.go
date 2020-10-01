package main

import (
	"fmt"
	"encoding/json"
	"os"
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

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// creates a struct based on the coordinate (c), and passes it into the
// json.Marshal function to format into a byte array. If it doesn't work, it
// returns an error instead.
func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DMS string `json: "dms"`
		DEC float64 `json: "decimal"`
		D float64 `json: "degrees"`
		M float64 `json: "minutes"`
		S float64 `json: "seconds"`
		H rune `json: "hemisphere"`
	}{
		DMS: c.String(),
		DEC: c.toDecimal(),
		D: c.d,
		M: c.m,
		S: c.s,
		H: c.h,
	})
}

type location struct {
	name string
	lat, long coordinate
}

func (l location) printCoordJSON() {
	fmt.Printf("%v\n", l.name)
	laBytes, err1 := json.MarshalIndent(l.lat, "", "	")
	exitOnError(err1)
	fmt.Printf("Latitude:\n%v\n", string(laBytes))
	loBytes, err2 := json.MarshalIndent(l.long, "", "	")
	exitOnError(err2)
	fmt.Printf("Longitude:\n%v\n", string(loBytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	london := location{
		name: "Paris",
		lat: coordinate{d: 48.0, m: 51.0, h: 'N',},
		long: coordinate{d: 2.0, m: 21.0, h: 'E',},
	}

	london.printCoordJSON()
}