package main

import (
	"fmt"
	"errors"
)

type sudokuError struct {
	err string
	x, y int
	v int8
}

// Use this when there's an error regarding a specific coordinate or value
func NewError(e string, y,x int, v int8) *sudokuError {
	return &sudokuError{err: e, y: y, x: x, v: v,}
}

func (sg sudokuError) Error() string {
	if sg.x > -1 {
		return fmt.Sprintf("Error: %v, Coords: [%v, %v], Value: %v\n", sg.err, sg.x, sg.y, sg.v)
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

func isValidCoord(y,x int) bool {
	return x >= 0 && x < 9 && y >= 0 && y < 9
}

func isValidValue(v int8) bool {
	return v >= 0 && v < 10
}

func (sg *sudokuGrid) CanPlace(y,x int, v int8) *sudokuError {
	var errPt *sudokuError
	if isValidCoord(y,x) {
		if isValidValue(v) {
			if !sg.initState[y][x] {
				errPt = NewError("Cannot replace values from the initial puzzle state", y, x, v)
			}
		} else {
			errPt = NewError("Sudoku values must be between 0 and 9", y, x, v)
		}
	} else {
		errPt = NewError("Out of range (x and y values need to be between 0 and 8)", y, x, v)
	}
	return errPt
}

func (sg *sudokuGrid) PrintCurrState() {
	for y := range sg.currState {
		fmt.Printf("%v\n", sg.currState[y])
	}
}

func (sg *sudokuGrid) PlaceNum(y,x int, v int8) *sudokuError {
	err := sg.CanPlace(y, x, v)
	if err == nil {
		sg.currState[y][x] = v
	}
	return err
}

// This function is slightly unnecessary - it's just shorthand for PlaceNum(y, x, 0).
func (sg *sudokuGrid) RemoveNum(y,x int) *sudokuError {
	var errPt *sudokuError
	if sg.initState[y][x] {
		sg.currState[y][x] = 0
	} else {
		errPt = NewError("Cannot replace values from the initial puzzle state", y, x, 0)
	}
	return errPt
}

func (sg *sudokuGrid) CheckCoord(y,x int) (bool, error) {
	if isValidCoord(y,x) {
		if sg.initState[y][x] {
			return true, nil
		}
		if sg.currState[y][x] == 0 {
			return false, nil
		}
		for h := range sg.currState {
			if sg.currState[h][x] == sg.currState[y][x] {
				return false, nil
			}
		}
		for w := range sg.currState[y] {
			if sg.currState[y][w] == sg.currState[y][x] {
				return false, nil
			}
		}
		var beginH int = y - (y % 3)
		var beginW int = x - (x % 3)
		for h := beginH; h < beginH + 3; h++ {
			for w := beginW; w < beginW + 3; w++ {
				if sg.currState[h][w] == sg.currState[y][x] {
					return false, nil
				}
			}
		}
	} else {
		return false, errors.New(fmt.Sprintf("Index out of range ([%v][%v])\n", y, x))
	}
	return true, nil
}

func HandleError(err error) {
	if e, ok := err.(*sudokuError); ok {
		fmt.Printf(e.Error())
	} else {
		fmt.Printf(err.Error())
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

	fmt.Println("Here is our initial puzzle state:")
	sg.PrintCurrState()	

	fmt.Println("Testing error handling with out of range values:")
	if err := sg.PlaceNum(9,9,1); err != nil {
		HandleError(err)
	} else {
		fmt.Println("Something went wrong while testing out of range errors in PlaceNum")
	}

	fmt.Println("Testing error handling with invalid sudoku puzzle value:")
	if err := sg.PlaceNum(0,2,1); err != nil {
		HandleError(err)
	} else {
		fmt.Println("Something went wrong while testing invalid sudoku input in PlaceNum")
	}

	fmt.Println("Testing error handling with attempted init puzzle state overwrites:")
	if err := sg.PlaceNum(0,0,1); err != nil {
		HandleError(err)
	} else {
		fmt.Println("Something went wrong while testing init state overwrite errors in PlaceNum")
	}

	fmt.Println("If the state has changed, something is wrong:")
	sg.PrintCurrState()

	fmt.Println("Testing the Error handling in CheckCoord method (should return index out of range):")
	if _, err := sg.CheckCoord(9,9); err != nil {
		HandleError(err)
	} else {
		fmt.Println("Something went wrong while testing out of range errors in CheckCoord")
	}

	var pFunc = func(p bool, out string) {
		fmt.Printf("Testing " + out + ": Test ")
		if p {
			fmt.Printf("Success\n")
		} else {
			fmt.Printf("Fail\n")
		}
	}
	p1, _ := sg.CheckCoord(0,0)
	pFunc(p1, "if preexisting sudoku values pass CheckCoord")
	p2, _ := sg.CheckCoord(0,2)
	pFunc(!p2, "if empty values in puzzle fail CheckCoord")
	sg.PlaceNum(0,2,7)
	p3, _ := sg.CheckCoord(0,2)
	pFunc(!p3, "if wrong value placed in puzzle failed to pass CheckCoord (x axis)")
	sg.PlaceNum(1,1,6)
	p4, _ := sg.CheckCoord(0,2)
	pFunc(!p4, "if wrong value placed in puzzle failed to pass CheckCoord (y axis)")
	sg.RemoveNum(0,2)
	sg.PlaceNum(1,1,4)
	p5, _ := sg.CheckCoord(1,1)
	pFunc(!p5, "if wrong value placed in puzzle failed to pass CheckCoord (3 x 3)")
}