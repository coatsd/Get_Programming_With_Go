package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type activity interface {
	move() string
	eat(food) string
}

type food string

type marsAnimal struct {
	kind, name string
}

func (ma marsAnimal) String() string {
	return fmt.Sprintf("%v the %v", ma.name, ma.kind)
}

func (ma marsAnimal) move() string {
	dist := rand.Intn(5) + 1
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

type report struct {
	hour int
	action string
}

func (r report) String() string {
	hour, ampm := convertHour(r.hour)
	return fmt.Sprintf("At %v %v, %v", hour, ampm, r.action)
}

func convertHour(hour int) (int, string) {
	if hour > 12 {
		return hour - 12, "PM"
	} else if hour == 0 {
		return hour + 12, "AM"
	} else if hour == 12 {
		return hour, "PM"
	}
	return hour, "AM"
}

func generateReports(start, runTime int, a []activity, f []food) []report {
	reports := make([]report, 0, 1)
	for i := 0; i < runTime; i++ {
		var r report
		start %= 24
		r.hour = start
		animal := a[rand.Intn(len(a))]
		if start < 7 || start > 22 {
			r.action = "Everyone is sound asleep."
		} else {
			if rand.Intn(2) == 1 {
				r.action = animal.move()
			} else {
				r.action = animal.eat(f[rand.Intn(len(f))])
			}	
		}
		reports = append(reports, r)
		start++
	}
	return reports
}

func main() {
	const startHour int = 8
	const runTime int = 72
	b := newBoarst("Molly")
	g := newGrog("Greg")
	p := newPelpine("Sarah")

	animals := []activity{b, g, p,}
	foods := []food{"red carrot", "marswheat", "space pear"}

	fmt.Printf("Created %v, %v, and %v", b, g, p)

	reports := generateReports(startHour, runTime, animals, foods)

	for _, r := range reports {
		fmt.Println(r)
	}
}