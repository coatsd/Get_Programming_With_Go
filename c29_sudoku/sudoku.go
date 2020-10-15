package main

import (
	"fmt"
	"error"
)

type sudokuError struct {
	err string
	x, y int
}

func NewError(e string, x,y int) sudokuError {
	return sudokuError{err: e, x: x, y: y,}
}

func (sg *sudokuError) Error() string {
	if sg.x > -1 {
		return fmt.Sprintf("Error: %v, Coords: %v, %v", sg.Err, sg.x, sg.y)
	}
	return fmt.Sprintf("Error: %v", sg.err)
}

type sudokuGrid [9][9]int8

func NewSudoku(grid [9][9]int8) sudokuGrid {
	var sg sudokuGrid = grid
	return sg
}

func (sg *sudokuGrid) ValidateGrid() bool {

}

func (sg *sudokuGrid) PlaceNum(num int8) error {

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