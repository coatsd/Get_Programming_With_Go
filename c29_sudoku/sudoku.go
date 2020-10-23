package main

import (
	"fmt"
	"error"
)

type sudokuError struct {
	err string
	x, y int
	v int8
}

func NewError(e string, y,x int, v int8) sudokuError {
	return sudokuError{err: e, y: y, x: x, v: v,}
}

func (sg *sudokuError) Error() string {
	if sg.x > -1 {
		return fmt.Sprintf("Error: %v, Coords: %v, %v, Value: %v", sg.Err, sg.x, sg.y, sg.v)
	}
	return fmt.Sprintf("Error: %v", sg.err)
}

type sudokuGrid struct {
	initState, currState [9][9]int8
}

func NewSudoku(grid [9][9]int8) sudokuGrid {
	return sg{initState: grid, currState: grid,}
}

func (sg *sudokuGrid) CanPlace(y,x int, v int8) (bool, sudokuError) {
	if x >= 0 && x < 9 && y >= 0 && y < 9 {
		if sg.initState[y][x] == 0 {
			return true, nil
		} else {
			return false, NewError("Cannot replace values from the initial puzzle state", y, x, v)
		}
	} else {
		return false, NewError("Out of range (x and y values need to be between 0 and 8)", y, x, v)
	}
}

func (sg *sudokuGrid) ValidateGrid() (bool, sudokuError) {
	
}

func (sg *sudokuGrid) PlaceNum(y,x int, v int8) error {
	can, err := CanPlace(y, x, v)
	if can {
		sg.currState[y][x] = v
	}
	return err
}

func (sg *sudokuGrid) RemoveNum(num int8) error {

}

func HandleError(err error) {
	if e, ok := err.(sudokuError); ok {

	}
}

func main() {
	s := NewSudoku([9][9]int8{
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
}