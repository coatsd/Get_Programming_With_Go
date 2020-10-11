package main

import (
	"fmt"
	"math/rand"
	"time"
)

type universe [][]bool

// Creates a new instance of a universe - every cell is initialized to false
func newUniverse(height int, width int) universe {
	u := make([][]bool, height)
	for w := range u {
		u[w] = make([]bool, width)
	}
	return u
}

// prints out a universe's 2d array, using 1 for true and 0 for false
func (u universe) showUniverse() {
	c0 := "0"
	c1 := "1"
	for _, h := range u {
		for _, w := range h {
			if w {
				fmt.Printf(" %v", c1)	
			} else {
				fmt.Printf(" %v", c0)
			}
		}
		fmt.Printf("\n")
	}
}

// calculates the percentage of cells that are currently alive
func (u universe) calcPopPercent() float64 {
	var totalCells int
	var liveCells int
	for _, h := range u	{
		totalCells += len(h)
		for w := range h {
			if h[w] == true {
				liveCells++
			}
		}
	}
	return float64(liveCells) / float64(totalCells) * 100
}

// prints a universe's state and percentage of living cells
func (u universe) printUniverseInfo() {
	u.showUniverse()
	fmt.Printf("living cell percentage: %v%%\n", u.calcPopPercent())
}

// seeds a universe, with a 25% chance for each cell to become alive
func (u universe) seedUniverse() {
	rand.Seed(time.Now().UnixNano())
	for _, h := range u	{
		for w, _ := range h {
			h[w] = rand.Intn(4) == 3
		}
	}
}

// If x or y are out of range, it loops the coordinate to the other side
func (u universe) treatCoord(y, x int) (int, int) {
	if y >= len(u) {
		y = y % len(u)
	} else if y < 0 {
		y = (y % -len(u)) + len(u)
	}
	if x >= len(u[y]) {
		x = x % len(u[y])
	} else if x < 0 {
		x = (x % -len(u[y])) + len(u[y])
	}
	return y, x
}

// tests a cell to see if it is alive (set to true)
func (u universe) isAlive(y, x int) bool {
	return u[y][x]
}

// calculates how many cells around a coordinate has a living cell
func (u universe) neighbors(y, x int) int {
	var n int = 0
	var ty, tx int
	for h := -1; h < 2; h++ {
		for w := -1; w < 2; w++ {
			if h != 0 || w != 0 {
				ty, tx = u.treatCoord(y + h, x + w)
				if u[ty][tx] {
					n++
				}
			}
		}
	}
	return n
}

// calculates if a cell will continue to live, or if a cell will start living
func (u universe) willLive(y, x int) bool {
	var n int = u.neighbors(y, x)
	if u.isAlive(y, x) {
		if n < 2 || n > 3 {
			return false
		}
		return true
	}
	if n == 3 {
		return true
	}
	return false
}

func main() {
	const (
		height = 10
		width = 10
	)

	initU := newUniverse(height, width)
	initU.seedUniverse()
	fmt.Printf("Initial Universe:\n")
	initU.printUniverseInfo()
	nextU := newUniverse(height, width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nextU[y][x] = initU.willLive(y, x)
		}
	}
	fmt.Printf("\nNext Generation:\n")
	nextU.printUniverseInfo()
	initU = nextU
}