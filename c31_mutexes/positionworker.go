package main

import (
	"fmt"
	"image"
	"time"
)

func posWorker() {
	pos := image.Point{X: 10, Y: 10}
	dir := image.Point{X: 1, Y: 0}
	delayMult := 1000
	next := time.After(time.Duration(delayMult) * time.Millisecond)
	for {
		if delayMult > 3000 {
			break
		}
		select {
		case <-next:
			pos = pos.Add(dir)
			fmt.Println("current position is : ", pos)
			delayMult += 500
			next = time.After(time.Duration(delayMult) * time.Millisecond)
		}
	}
}

func main() {
	posWorker()
}