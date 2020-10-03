package main

import (
	"fmt"
)

type turtle struct {
	x, y int
}

func (t turtle) String() string {
	return fmt.Sprintf("X: %v, Y: %v", t.x, t.y)
}

func (t *turtle) move(x, y int) {
	t.x += x
	t.y += y
}

func main() {
	t := turtle{2, 2}
	fmt.Println(t)
	t.move(1, 3)
	fmt.Println(t)
}