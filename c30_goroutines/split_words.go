package main

import (
	"fmt"
	"strings"
)

type stringChan struct {
	c chan string
	w []string
}

func newStringChan(s string) *stringChan {
	words := strings.Fields(s)
	c := make(chan string, len(words))
	return &stringChan{c, words}
}

func (sc *stringChan) ReadAll() {
	for _, v := range sc.w {
		sc.c <- v
	}
	close(sc.c)
}

func printChan(c chan string) {
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println(v)
	}
}

func makeAndSend(s string, scChan chan *stringChan, done chan bool) {
	scChan <- newStringChan(s)
	done <- true
}

func source(scChan chan *stringChan) {
	done := make(chan bool)
	counter := 0
	stringVals := []string{"Hello World", "Gopher Power", "Good Day Sir"}
	for _, v := range stringVals {
		go makeAndSend(v, scChan, done)
	}
	for _, v := range stringVals {
		if <-done {
			counter++
		}
		fmt.Println("Made chan for: ", v)
	}
	if counter == len(stringVals) {
		close(scChan)	
	}
}

func readSCs(us chan *stringChan) {
	for {
		sc, ok := <-us
		if !ok {
			break
		}
		go sc.ReadAll()
		printChan(sc.c)
	}
}

func main() {
	scChan := make(chan *stringChan)
	go source(scChan)
	readSCs(scChan)
}