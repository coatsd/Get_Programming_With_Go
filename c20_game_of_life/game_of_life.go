package main

import (
	"fmt"
	"math/rand"
	"time"
)

type universe [][]bool

// Creates a new instance of a universe - every cell is initialized to false
func newUniverse(width int, height int) universe {
	w := make([][]bool, width)
	for h, _ := range w {
		w[h] = make([]bool, height)
	}
	return w
}

// prints out a universe's 2d array, using 1 for true and 0 for false
func (u universe) showUniverse() {
	c0 := "0"
	c1 := "1"
	for _, w := range u {
		for _, h := range w {
			if h {
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
	for _, w := range u	{
		totalCells += len(w)
		for h, _ := range w {
			if w[h] == true {
				liveCells++
			}
		}
	}
	return float64(liveCells) / float64(totalCells) * 100
}

// seeds a universe, with a 25% chance for each cell to become alive
func (u universe) seedUniverse() {
	rand.Seed(time.Now().UnixNano())
	for _, w := range u	{
		for h, _ := range w {
			w[h] = rand.Intn(4) == 3
		}
	}
}

// If x or y are out of range, it loops the coordinate to the beginning
func (u universe) treatCoord(x int, y int) (int, int) {
	if x >= len(u) {
		x = x % len(u)
	} else if x < 0 {
		x = (x % -len(u)) + len(u)
	}
	if y >= len(u[x]) {
		y = y % len(u[x])
	} else if y < 0 {
		y = (y % -len(u[x])) + len(u[x])
	}
	return x, y
}

// tests a cell to see if it is alive (set to true)
func (u universe) isAlive(x int, y int) bool {
	return u[x][y]
}

// calculates how many cells around a coordinate has a living cell
func (u universe) neighbors(x int, y int) int {
	var n int = 0
	for w := -1; w < 2; w++ {
		for h := -1; h < 2; h++ {
			if w != 0 || h != 0 {
				x, y = u.treatCoord(x + w, y + h)
				if u[x][y] {
					n++
				}
			}
		}
	}
	return n
}

// prints a universe's state and percentage of living cells
func (u universe) printUniverseInfo() {
	u.showUniverse()
	fmt.Printf("living cell percentage: %v%%\n", u.calcPopPercent())
}

// calculates if a cell will continue to live, or if a cell will start living
func (u universe) willLive(x int, y int) bool {
	var n int = u.neighbors(x, y)
	if u.isAlive(x, y) {
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
		width = 20
		height = 10
	)

	initU := newUniverse(width, height)
	initU.seedUniverse()
	fmt.Printf("Initial Universe:\n")
	initU.printUniverseInfo()
	nextU := newUniverse(width, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			nextU[x][y] = initU.willLive(x, y)
		}
	}
	fmt.Printf("\nNext Generation:\n")
	nextU.printUniverseInfo()
}