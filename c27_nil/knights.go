package main

import (
	"fmt"
)

type item struct {
	name string
	weight float64
}

type inventory []*item

type character struct {
	name string
	i inventory
	rightH, leftH *item
}

func (owner *character) ReleaseItem(index int) *item {
	var i *item
	if index < len(owner.i) && index >= 0 {
		i, owner.i = owner.i[index], append(owner.i[:index], owner.i[index+1:]...)
	}
	return i
}

func (c *character) GetItem(i *item) item {
	c.i = append(c.i, i)
	return *i
}

func (c *character) EquipRH(i *item) item {
	if c.rightH == nil {
		c.rightH = i
	} else {
		c.i, c.rightH = append(c.i, c.rightH), i
	}
	return *c.rightH
}

func (c *character) EquipLH(i *item) item {
	if c.leftH == nil {
		c.leftH = i
	} else {
		c.i, c.leftH = append(c.i, c.leftH), i
	}
	return *c.leftH
}

func main() {
	k := character{name: "knight"}
	a := character{name: "Arthur"}
	i := item{name: "sword", weight: 2.5}

	fmt.Printf("Arthur picked up: %v\n", a.GetItem(&i).name)
	fmt.Printf("Arthur gave %v to the knight.\n", k.GetItem(a.ReleaseItem(0)).name)
	fmt.Printf("Knight's Left hand Item: %v\n", k.EquipLH(k.ReleaseItem(0)).name)
}