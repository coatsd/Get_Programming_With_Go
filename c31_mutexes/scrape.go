package main

import (
	"fmt"
	"sync"
)

type Visited struct {
	mu sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	c := v.visited[url]
	c++
	v.visited[url] = c
	return c
}

func main() {
	
}