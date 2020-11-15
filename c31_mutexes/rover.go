package main

import (
	"fmt"
	"time"
	"image"
)

type command int

const (
	stop = command(0)
	start = command(1)
	left = command(2)
	right = command(3)
)

type RoverDriver struct {
	c chan command
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{c: make(chan command),}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0,}
	dir := image.Point{X: 0, Y: 0,}
	dirState := image.Point{X: 1, Y: 0,}
	updateInterval := 500 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.c:
			switch c {
			case stop:
				dirState = dir
				dir = image.Point{X: 0, Y: 0,}
				fmt.Println("Rover stopped")
			case start:
				dir = dirState
				fmt.Println("Rover started")
			case right:
				dir = image.Point{X: -dir.Y, Y: dir.X,}
				fmt.Println("turned right") 
			case left:
				dir = image.Point{X: dir.Y, Y: -dir.X,}
				fmt.Println("turned left")
			}
		case <-nextMove:
			pos = pos.Add(dir)
			fmt.Printf("moved to %v\n", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) stop() {
	r.c <- stop
}

func (r *RoverDriver) start() {
	r.c <- start
}

func (r *RoverDriver) left() {
	r.c <- left
}

func (r *RoverDriver) right() {
	r.c <- right
}

func main() {
	r := NewRoverDriver()
	r.start()
	time.Sleep(3 * time.Second)
	r.left()
	time.Sleep(3 * time.Second)
	r.right()
	time.Sleep(3 * time.Second)
	r.stop()
	time.Sleep(3 * time.Second)
	r.start()
	time.Sleep(3 * time.Second)
}