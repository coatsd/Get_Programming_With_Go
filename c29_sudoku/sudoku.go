package main

import (
	"fmt"
	"os"
)

type sudokuError struct {
	err string
	x, y int
	v int8
}

func NewError(e string, y,x int, v int8) *sudokuError {
	return &sudokuError{err: e, y: y, x: x, v: v,}
}

func (sg sudokuError) Error() string {
	if sg.x > -1 {
		return fmt.Sprintf("Error: %v, Coords: %v, %v, Value: %v\n", sg.err, sg.x, sg.y, sg.v)
	}
	return fmt.Sprintf("Error: %v\n", sg.err)
}

type sudokuGrid struct {
	initState [9][9]bool
	currState [9][9]int8
}

func NewSudoku(grid [9][9]int8) sudokuGrid {
	var init [9][9]bool
	for h := range grid {
		for w := range grid[h] {
			init[h][w] = grid[h][w] == 0
		}
	}
	return sudokuGrid{initState: init, currState: grid,}
}

func (sg *sudokuGrid) CanPlace(y,x int, v int8) *sudokuError {
	var errPt *sudokuError
	if x >= 0 && x < 9 && y >= 0 && y < 9 {
		if !sg.initState[y][x] {
			errPt = NewError("Cannot replace values from the initial puzzle state", y, x, v)
		}
	} else {
		errPt = NewError("Out of range (x and y values need to be between 0 and 8)", y, x, v)
	}
	return errPt
}

func (sg *sudokuGrid) PlaceNum(y,x int, v int8) *sudokuError {
	err := sg.CanPlace(y, x, v)
	if err != nil {
		sg.currState[y][x] = v
	}
	return err
}

func (sg *sudokuGrid) RemoveNum(y,x int) *sudokuError {
	var errPt *sudokuError
	if sg.initState[y][x] {
		sg.currState[y][x] = 0
	} else {
		errPt = NewError("Cannot replace values from the initial puzzle state", y, x, 0)
	}
	return errPt
}

func HandleError(err error) {
	if e, ok := err.(*sudokuError); ok {
		fmt.Printf(e.Error())
	} else {
		fmt.Printf(e.Error())
		os.Exit(1)
	}
}

func main() {
	sg := NewSudoku([9][9]int8{
		{5,3,0,0,7,0,0,0,0},
		{6,0,0,1,9,5,0,0,0},
		{0,9,8,0,0,0,0,6,0},
		{8,0,0,0,6,0,0,0,3},
		{4,0,0,8,0,3,0,0,1},
		{7,0,0,0,2,0,0,0,6},
		{0,6,0,0,0,0,2,8,0},
		{0,0,0,4,1,9,0,0,5},
		{0,0,0,0,8,0,0,7,9},
	})

	HandleError(sg.CanPlace(9,9,1))
}