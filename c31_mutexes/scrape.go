package main

import (
	"fmt"
	"sync"
)

type Visited struct {
	mu sync.Mutex
	visited map[string]int
}
func makeVisited() Visited {
	return Visited{visited: make(map[string]int)}
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	c := v.visited[url]
	c++
	v.visited[url] = c
	fmt.Println(url)
	return c
}

func main() {
	v := makeVisited()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		v.VisitLink("https://www.facebook.com")
	}()
	go func() {
		defer wg.Done()
		v.VisitLink("https://www.google.com")
	}()
	wg.Wait()
}