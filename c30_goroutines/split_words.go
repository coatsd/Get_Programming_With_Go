package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
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

func source(scChan chan *stringChan) {
	var wg sync.WaitGroup
	stringVals := []string{"Hello World", "Gopher Power", "Good Day Sir"}
	wg.Add(len(stringVals))
	for _, v := range stringVals {
		go func(s string, scChan chan *stringChan, wg *sync.WaitGroup) {
			defer wg.Done()
			scChan <- newStringChan(s)	
		}(v, scChan, &wg)
	}
	wg.Wait()
	close(scChan)
}

func readSCs(us chan *stringChan) {
	for {
		sc, ok := <-us 
		if !ok {
			break
		}
		var wg sync.WaitGroup
		wg.Add(cap(sc.c))
		go sc.ReadAll()
		go func(c chan string, wg *sync.WaitGroup) {
			for {
				v, ok := <-c
				if !ok {
					break
				}
				defer wg.Done()
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
				fmt.Println(v)
			}
		}(sc.c, &wg)
		wg.Wait()
	}
}

func main() {
	scChan := make(chan *stringChan)
	go source(scChan)
	readSCs(scChan)
}