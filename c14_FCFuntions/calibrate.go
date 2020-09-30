package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	offset := kelvin(5)
	s1 := calibrate(realSensor, offset)
	s2 := calibrate(fakeSensor, offset)
	fmt.Println(s1())
	fmt.Println(s2())
	/*
	Doesn't change value. When functions are generated, they don't pass a
	reference to offset - it passes a copy of offset. This could be fixed by
	creating the calibrate function here in main, remove the offset parameter,
	and reference offset directly in the body of calibrate. Either that, or
	perhaps a pointer to the value would work too.
	*/
	offset++
	fmt.Println(s1())
	fmt.Println(s2())
}