package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type activities interface {
	move() string
	eat() string
}

type food struct {
	name string
}

func (f food) String() string {
	return f.name
}

type marsAnimal struct {
	kind, name string
}

func (ma marsAnimal) String() string {
	return fmt.Sprintf("%v the %v", ma.name, ma.kind)
}

func (ma marsAnimal) move() string {
	dist := rand.Intn(5)
	return fmt.Sprintf("%v walked %v meters.", ma, dist)
}

func (ma marsAnimal) eat(f food) string {
	return fmt.Sprintf("%v ate %v", ma, f)
}

type boarst struct {
	marsAnimal
}

func newBoarst(name string) boarst {
	b := boarst {
		marsAnimal{"", name},
	}
	b.kind = reflect.TypeOf(b).Name()
	return b
}

type grog struct {
	marsAnimal
}

func newGrog(name string) grog {
	g := grog{
		marsAnimal{"", name},
	}
	g.kind = reflect.TypeOf(g).Name()
	return g
}

type pelpine struct {
	marsAnimal
}

func newPelpine(name string) pelpine {
	p := pelpine{
		marsAnimal{"", name},
	}
	p.kind = reflect.TypeOf(p).Name()
	return p
}

func main() {
	b := newBoarst("Molly")
	g := newGrog("Greg")
	p := newPelpine("Sarah")

	fmt.Printf("Created %v, %v, and %v", b, g, p)
}